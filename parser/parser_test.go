package parser

import (
	"fmt"
	"testing"
)

func TestParse01(t *testing.T) {
	code := "a + b.foo()"
	tokens, err := Tokenize(code)
	if err != nil {
		t.Fatal(err)
	}
	node, err := Parse(tokens)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(node.type_)
}

func TestParse02(t *testing.T) {
	code := "a + b.foo(-)"
	tokens, err := Tokenize(code)
	if err != nil {
		t.Fatal(err)
	}
	node, err := Parse(tokens)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(node.type_)
}
