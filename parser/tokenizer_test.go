package parser

import (
	"fmt"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tokens, err := Tokenize("1 + b")
	if err != nil {
		t.Fatal(err)
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
}
