package gui

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	root := Div(NewOpt(),
		Text(NewOpt(), "'hello'"),
	)
	model, err := GenHtml("", root)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(model)
}
