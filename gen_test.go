package gui

import (
	"fmt"
	"github.com/lincaiyong/gui/com/root"
	"github.com/lincaiyong/gui/com/text"
	"testing"
)

func TestGen(t *testing.T) {
	comp := root.Root(text.Text("'xx'"))
	root.AddProp("test", "1")
	html, err := MakeHtml("", comp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(html)
}
