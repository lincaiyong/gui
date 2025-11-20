package gui

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/lincaiyong/gui/parser"
	"io/fs"
	"reflect"
	"strings"
)

func GenModel(root *Element) (string, error) {
	pr := NewPrinter()
	err := genModel(root, 0, pr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("const root = %s;", strings.TrimRight(pr.Code(), ",")), nil
}

func GenHtml(title string, root *Element) (string, error) {
	model, err := GenModel(root)
	if err != nil {
		return "", fmt.Errorf("fail to gen model: %w", err)
	}
	pr := NewPrinter().Push().Push()
	err = copyJsCode(pr)
	if err != nil {
		return "", fmt.Errorf("fail to copy js code: %w", err)
	}
	pr.Put(model)
	code := pr.Code()
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title>{{title}}</title>
	<style>
		svg svg {
		  width: 100%;
		  height: 100%;
		  display: block;
		}
	</style>
    <link rel="stylesheet" href="res/vs/editor/editor.main.css" />
</head>
<body>
    <script src="res/vs/loader.js"></script>
    <script>
		{{code}}
        require.config({paths: {'vs': 'res/vs'}});
        require(['vs/editor/editor.main'], () => g.createAll(document.body, root));
    </script>
</body>
</html>`
	html = strings.ReplaceAll(html, "{{title}}", title)
	html = strings.ReplaceAll(html, "{{code}}", code)
	return html, nil
}

//go:embed js/*
var jsEmbed embed.FS

func copyJsCode(pr *Printer) error {
	items, err := fs.ReadDir(jsEmbed, ".")
	if err != nil {
		return err
	}
	for _, item := range items {
		var b []byte
		b, err = fs.ReadFile(jsEmbed, item.Name())
		if err != nil {
			return err
		}
		pr.Put(string(b))
	}
	return nil
}

func genModel(ele *Element, depth int, pr *Printer) error {
	t := reflect.ValueOf(ele).Type().Elem()
	pr.Put("{").Push()
	{
		s := t.Name()
		if s == "Component" {
			s = t.String()
			s = s[:strings.Index(s, ".")]
			s = pascalCase(s)
		}
		compName := ele.Name()
		if compName == "" {
			compName = ele.Tag()
		}
		pr.Put("tag: '%s',", ele.Tag())
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
		for k, v := range ele.Props() {
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
		children := ele.Children()
		slots := ele.Slots()
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
				err := genModel(tmp, childrenDepth, pr)
				if err != nil {
					return err
				}
			}
			pr.Pop().Put("],")
		}
		if len(slots) > 0 {
			pr.Put("slot: [").Push()
			for _, tmp := range slots {
				err := genModel(tmp, depth+1, pr)
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
