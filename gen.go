package gui

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/parser"
	"github.com/lincaiyong/gui/printer"
	"github.com/lincaiyong/gui/utils"
	"github.com/lincaiyong/gui/visit"
	"github.com/lincaiyong/log"
	"net/http"
	"reflect"
	"sort"
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
		pr.Put("tag: '%s',", comp.Tag())
		pr.Put("id: '%s',", compName)
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
	model, err := buildPageModel(page)
	if err != nil {
		return "", err
	}
	s := []string{guiJs, model}
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
