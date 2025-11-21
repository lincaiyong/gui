package gui

import (
	"fmt"
	"testing"
)

func TestCom(t *testing.T) {
	root := Div(NewOpt())
	model, err := GenModel(root)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(model)
}
