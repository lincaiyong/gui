package gui

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/com/root"
	"github.com/lincaiyong/gui/js"
	"github.com/lincaiyong/gui/parser"
	"github.com/lincaiyong/gui/printer"
	"github.com/lincaiyong/gui/utils"
	"github.com/lincaiyong/gui/visit"
	"github.com/lincaiyong/log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var baseUrl = "http://127.0.0.1:9123"

func SetBaseUrl(url string) {
	baseUrl = url
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getAllInstances(comp com.Component, comps []com.Component, thisComp com.Component, selfIndex []int) []com.Component {
	struct_ := reflect.TypeOf(comp).Elem()
	name := struct_.Name()
	if name == "Component" {
		name = struct_.String()
		name = name[:strings.Index(name, ".")]
		name = utils.PascalCase(name)
	}
	info := comp.ExtraInfo()
	info.SetName(name)
	if thisComp == nil {
		thisComp = comp
	}
	info.SetThisComponent(thisComp)
	info.SetSelfIndex(selfIndex)
	comps = append(comps, comp)
	for i, tmp := range comp.Slots() {
		comps = getAllInstances(tmp, comps, thisComp, append(selfIndex, i))
	}
	for i, tmp := range comp.Children() {
		comps = getAllInstances(tmp, comps, comp, []int{i})
	}
	return comps
}

func readFunction(name, code string) (string, error) {
	if !strings.HasPrefix(code, fmt.Sprintf("function %s(", name)) {
		return "", fmt.Errorf("invalid function code: %s", code)
	}
	return code[9:], nil
}

func genClassCode(info *com.ExtraInfo, namedChildren map[string]map[string][]int, pr *printer.Printer) error {
	pr.Put("class %s extends Component {", info.Name()).Push()
	{
		pr.Put("constructor(parent, model) {").Push()
		{
			if len(info.Properties()) == 0 {
				pr.Put("model.properties = Object.assign({}, model.properties);")
			} else {
				pr.Put("model.properties = Object.assign({").Push()
				for _, property := range info.Properties() {
					pr.Put("%s: [e => %s, []],", property, info.GetDefaultValue(property))
				}
				pr.Pop().Put("}, model.properties);")
			}
			pr.Put("super(parent, model);")
		}
		pr.Pop().Put("}")
		for _, method := range info.Methods() {
			code, err := readFunction(method, info.GetBindJs(method))
			if err != nil {
				return err
			}
			pr.Put(code)
		}
		for _, method := range info.StaticMethods() {
			code, err := readFunction(method, info.GetBindJs(method))
			if err != nil {
				return err
			}
			pr.Put("static " + code)
		}
		for _, property := range info.Properties() {
			pr.Put("get %s() {", property).Push().
				Put("return this._properties.%s.value;", property).Pop().Put("}")
			pr.Put("set %s(v) {", property).Push().
				Put("this._properties.%s.value = v;", property).Pop().Put("}")
		}
		if m, ok := namedChildren[info.Name()]; ok {
			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				items := make([]string, 0, len(m[k]))
				for _, i := range m[k] {
					items = append(items, strconv.Itoa(i))
				}
				pr.Put("get %s() { return [%s].reduce((prev, i) => prev.children[i], this); }", k, strings.Join(items, ", "))
			}
		}
	}
	if info.Name() == "Root" {
		m := js.GetAll(info.Name())
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if strings.HasPrefix(m[k], "function ") {
				code := m[k][9:]
				pr.Put("static " + code)
			}
		}
	}
	pr.Pop().Put("}")
	return nil
}

func buildClasses(page com.Component) (string, error) {
	comps := getAllInstances(page, nil, nil, nil)
	namedChildren := make(map[string]map[string][]int)
	for _, comp := range comps {
		name := comp.Name()
		if name == "" {
			continue
		}
		if !strings.HasSuffix(name, "Ele") {
			return "", fmt.Errorf("invalid element name: %s", name)
		}
		thisComp := comp.ExtraInfo().ThisComponent()
		if thisComp != comp {
			thisCompName := thisComp.ExtraInfo().Name()
			if m, ok := namedChildren[thisCompName]; ok {
				m[comp.Name()] = comp.ExtraInfo().SelfIndex()
			} else {
				m = make(map[string][]int)
				m[comp.Name()] = comp.ExtraInfo().SelfIndex()
				namedChildren[thisCompName] = m
			}
		}
	}
	compMap := make(map[string]com.Component)
	for _, comp := range comps {
		name := comp.ExtraInfo().Name()
		if _, ok := compMap[name]; !ok {
			compMap[name] = comp
		}
	}
	keys := make([]string, 0, len(compMap))
	for n, comp := range compMap {
		keys = append(keys, n)
		info := comp.ExtraInfo()
		if info.Name() == "Root" {
			for k, v := range root.Props() {
				info.AddProperty(k)
				info.SetDefaultValue(k, v)
			}
		}

		struct_ := reflect.TypeOf(comp).Elem()
		for i := 0; i < struct_.NumField(); i++ {
			field := struct_.Field(i)
			if !field.Anonymous {
				tn := field.Type.Name()
				switch tn {
				case "Property":
					info.AddProperty(field.Name)
					v := field.Tag.Get("default")
					if v == "" {
						return "", fmt.Errorf("default value is required: %s", field.Name)
					}
					info.SetDefaultValue(field.Name, v)
				case "Method":
					if field.Tag.Get("static") == "true" {
						info.AddStaticMethod(field.Name)
					} else {
						info.AddMethod(field.Name)
					}
					code := js.Get(n, field.Name)
					if code == "" {
						return "", fmt.Errorf("fail to bind method: %s", field.Name)
					}
					info.SetBindJs(field.Name, code)
				}
			}
		}
	}
	sort.Strings(keys)
	pr := printer.NewPrinter()
	for _, k := range keys {
		err := genClassCode(compMap[k].ExtraInfo(), namedChildren, pr)
		if err != nil {
			return "", err
		}
	}
	return pr.Code(), nil
}

func convertExpr(s string) (string, string, error) {
	tokens, err := parser.Tokenize(s)
	if err != nil {
		return s, "[]", nil
	}
	node, err := parser.Parse(tokens)
	if err != nil {
		return "", "", err
	}
	s1, s2, err := visit.Visit(node)
	if err != nil {
		return "", "", err
	}
	return s1, s2, nil
}

func buildModel(comp com.Component, depth int, pr *printer.Printer) error {
	t := reflect.ValueOf(comp).Type().Elem()
	pr.Put("{").Push()
	{
		s := t.Name()
		if s == "Component" {
			s = t.String()
			s = s[:strings.Index(s, ".")]
			s = utils.PascalCase(s)
		}
		compName := comp.Name()
		if compName == "" {
			compName = comp.Tag()
		}
		pr.Put("Component: %s,", s)
		pr.Put("tag: '%s',", comp.Tag())
		pr.Put("name: '%s',", compName)
		pr.Put("depth: %d,", depth)
		props := make(map[string]string)
		if depth > 0 {
			props["h"] = "parent.ch"
			props["v"] = "parent.v"
			props["w"] = "parent.cw"
			props["x"] = "-parent.scrollLeft"
			props["y"] = "-parent.scrollTop"
			props["zIndex"] = "parent.zIndex"
		}
		for k, v := range comp.Props() {
			props[k] = v
		}
		if len(props) == 0 {
			pr.Put("properties: {},")
		} else {
			pr.Put("properties: {").Push()
			for _, k := range sortedKeys(props) {
				v1, v2, err := convertExpr(props[k])
				if err != nil {
					return err
				}
				pr.Put("%s: [e => %s, %s],", k, v1, v2)
			}
			pr.Pop().Put("},")
		}
		children := comp.Children()
		slots := comp.Slots()
		var childrenDepth int
		if s == "Div" {
			children = slots
			slots = nil
			childrenDepth = depth + 1
		} else if s == "Containeritem" { // 允许child通过this访问祖先
			childrenDepth = depth + 1
		} else {
			childrenDepth = 1
		}
		if len(children) > 0 {
			pr.Put("children: [").Push()
			for _, tmp := range children {
				err := buildModel(tmp, childrenDepth, pr)
				if err != nil {
					return err
				}
			}
			pr.Pop().Put("],")
		}
		if len(slots) > 0 {
			pr.Put("slot: [").Push()
			for _, tmp := range slots {
				err := buildModel(tmp, depth+1, pr)
				if err != nil {
					return err
				}
			}
			pr.Pop().Put("],")
		}
	}
	pr.Pop().Put("},")
	return nil
}

func buildPageModel(page com.Component) (string, error) {
	pr := printer.NewPrinter()
	err := buildModel(page, 0, pr)
	if err != nil {
		return "", err
	}
	code := pr.Code()
	code = strings.TrimRight(code, ",")
	return "g.model = " + code + ";", nil
}

//go:embed js/*.js
var jsEmbed embed.FS

func MakePage(c *gin.Context, title string, page *root.Component) {
	html, err := MakeHtml(title, page)
	if err != nil {
		log.ErrorLog("fail to make page: %v", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("fail to make page: %v", err))
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func MakeHtml(title string, page *root.Component) (string, error) {
	eventJs, _ := jsEmbed.ReadFile("js/_event.js")
	propertyJs, _ := jsEmbed.ReadFile("js/_property.js")
	scrollbarJs, _ := jsEmbed.ReadFile("js/_scrollbar.js")
	componentJs, _ := jsEmbed.ReadFile("js/_component.js")
	gJs, _ := jsEmbed.ReadFile("js/_g.js")
	classes, err := buildClasses(page)
	if err != nil {
		return "", err
	}
	model, err := buildPageModel(page)
	if err != nil {
		return "", err
	}
	s := []string{string(eventJs), string(propertyJs), string(scrollbarJs), string(componentJs), string(gJs), classes, model}
	ss := strings.Join(strings.Split(strings.Join(s, "\n"), "\n"), "\n        ")
	template := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title><ttt></title>
	<style>
		svg svg {
		  width: 100%;
		  height: 100%;
		  display: block;
		}
	</style>
    <link rel="stylesheet" href="<base_url>/res/vs/editor/editor.main.css" />
</head>
<body>
    <script src="<base_url>/res/vs/loader.js"></script>
    <script>
        <xxx>
        require.config({paths: {'vs': '<base_url>/res/vs'}});
        require(['vs/editor/editor.main'], () =>g.create());
    </script>
</body>
</html>`
	if !strings.Contains(ss, "monaco") {
		template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title><ttt></title>
	<style>
		svg svg {
		  width: 100%;
		  height: 100%;
		  display: block;
		}
	</style>
</head>
<body>
    <script>
        <xxx>
        g.create();
    </script>
</body>
</html>`
	}
	template = strings.ReplaceAll(template, "<ttt>", title)
	ss = strings.ReplaceAll(template, "<xxx>", ss)
	html := strings.ReplaceAll(ss, "<base_url>", baseUrl)
	return html, nil
}
