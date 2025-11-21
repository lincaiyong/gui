package gui

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
)

func GenModel(root *Element) (string, error) {
	walkTree(root, root, 0, nil)
	pr := NewPrinter()
	err := genModel(root, 0, pr)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(pr.Code(), ","), nil
}

func GenHtml(title string, root *Element) (string, error) {
	pr := NewPrinter().Push().Push()
	err := copyJsCode(pr)
	if err != nil {
		return "", fmt.Errorf("fail to copy js code: %w", err)
	}
	model, err := GenModel(root)
	if err != nil {
		return "", fmt.Errorf("fail to gen model: %w", err)
	}
	pr.Put("const root = %s;", model)
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
	items, err := fs.ReadDir(jsEmbed, "js")
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		var b []byte
		b, err = fs.ReadFile(jsEmbed, filepath.Join("js", item.Name()))
		if err != nil {
			return err
		}
		pr.Put(string(b))
	}
	return nil
}

func genModel(ele *Element, depth int, pr *Printer) error {
	pr.Put("{").Push()
	{
		pr.Put("tag: '%s',", ele.Tag())
		pr.Put("name: '%s',", ele.Name())
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
		for k, v := range ele.Properties() {
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
		if ele.Type() == ElementTypeContainer {
			child := children[2]
			children = []*Element{children[0], children[1]}
			tmpPr := NewPrinter()
			err := genModel(child, depth+1, tmpPr)
			if err != nil {
				return err
			}
			pr.Put("itemModel: %s", tmpPr.Code())
		}
		if len(children) > 0 {
			pr.Put("children: [").Push()
			for _, tmp := range children {
				if ele.IsLocalRoot() {
					depth = 0
				}
				err := genModel(tmp, depth+1, pr)
				if err != nil {
					return err
				}
			}
			pr.Pop().Put("],")
		}
		namedElements := make([]*Element, 0)
		for _, e := range ele.LocalElements() {
			if strings.HasSuffix(e.Name(), "Ele") {
				namedElements = append(namedElements, e)
			}
		}
		if len(namedElements) > 0 {
			pr.Put("named: {").Push()
			for _, child := range namedElements {
				idx := child.SelfIndex()
				items := make([]string, len(idx))
				for i, v := range idx {
					items[i] = strconv.Itoa(v)
				}
				pr.Put("%s: [%s],", child.Name(), strings.Join(items, ", "))
			}
			pr.Pop().Put("},")
		}
	}
	pr.Pop().Put("},")
	return nil
}

func walkTree(ele *Element, localRoot *Element, depth int, index []int) {
	ele.SetSelfIndex(index)
	localRoot.AddLocalElement(ele)
	if ele.IsLocalRoot() {
		localRoot = ele
		depth = 0
	}
	for i, child := range ele.Children() {
		walkTree(child, localRoot, depth+1, append(index, i))
	}
}
