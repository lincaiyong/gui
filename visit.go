package gui

import (
	"fmt"
	"github.com/lincaiyong/gui/parser"
	"sort"
	"strings"
)

var relations = map[string]bool{
	"":       true,
	"parent": true,
	"child":  true,
	"prev":   true,
	"next":   true,
	"local":  true,
	"root":   true,
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
	s1, s2, err := visit(node)
	if err != nil {
		return "", "", err
	}
	return s1, s2, nil
}

func visit(node *parser.Node) (string, string, error) {
	v := Visitor{
		m: map[string]struct{}{},
	}
	expr := v.visit(node)
	if v.err != nil {
		return "", "", v.err
	}
	var s []string
	for k := range v.m {
		s = append(s, k)
	}
	sort.Strings(s)
	return expr, fmt.Sprintf("[%s]", strings.Join(s, ", ")), nil
}

type Visitor struct {
	m   map[string]struct{}
	err error
}

func (v *Visitor) visit(node *parser.Node) string {
	switch node.Type() {
	case parser.NodeTypeBinary:
		lhs := v.visit(node.BinaryLhs())
		rhs := v.visit(node.BinaryRhs())
		return fmt.Sprintf("(%s %s %s)", lhs, node.Op(), rhs)
	case parser.NodeTypeSelector:
		var target string
		key := node.SelectorKey()
		if node.SelectorTarget() != nil {
			target = v.visit(node.SelectorTarget())
			if target == "e" {
				return fmt.Sprintf("function() { return e.%s(...arguments); }", key)
			}
		}
		if node.SelectorTarget() == nil {
			v.m[fmt.Sprintf("'.%s'", key)] = struct{}{}
			return fmt.Sprintf("e.%s", key)
		} else if node.SelectorTarget().Type() == parser.NodeTypeIdent && relations[node.SelectorTarget().Ident()] {
			v.m[fmt.Sprintf("'%s.%s'", node.SelectorTarget().Ident(), key)] = struct{}{}
			return fmt.Sprintf("%s.%s", target, key)
		} else {
			return fmt.Sprintf("%s.%s", target, key)
		}
	case parser.NodeTypeUnary:
		x := v.visit(node.UnaryTarget())
		return fmt.Sprintf("%s%s", node.Op(), x)
	case parser.NodeTypeArray:
		var items []string
		for _, n := range node.ArrayItems() {
			items = append(items, v.visit(n))
		}
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	case parser.NodeTypeCall:
		target := v.visit(node.Callee())
		var args []string
		for _, n := range node.Args() {
			args = append(args, v.visit(n))
		}
		return fmt.Sprintf("%s(%s)", target, strings.Join(args, ", "))
	case parser.NodeTypeTernary:
		con := v.visit(node.TernaryCondition())
		lhs := v.visit(node.TernaryLhs())
		rhs := v.visit(node.TernaryRhs())
		return fmt.Sprintf("%s ? %s : %s", con, lhs, rhs)
	case parser.NodeTypeString:
		return node.String()
	case parser.NodeTypeNumber:
		return node.Number()
	case parser.NodeTypeIdent:
		i := node.Ident()
		if relations[i] {
			i = fmt.Sprintf("e.%s", i)
		}
		return i
	default:
		v.err = fmt.Errorf("fail to visit unknown node type: %s", node.Type())
		return ""
	}
}
