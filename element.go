package gui

import (
	"fmt"
	"github.com/lincaiyong/log"
	"strings"
)

type ElementType string

const (
	ElementTypeDiv       ElementType = "div"
	ElementTypeText      ElementType = "text"
	ElementTypeInput     ElementType = "input"
	ElementTypeImg       ElementType = "img"
	ElementTypeSvg       ElementType = "svg"
	ElementTypeDivider   ElementType = "divider"
	ElementTypeBar       ElementType = "bar"
	ElementTypeScrollbar ElementType = "scrollbar"
	ElementTypeIframe    ElementType = "iframe"
	ElementTypeButton    ElementType = "button"
	ElementTypeCompare   ElementType = "compare"
	ElementTypeContainer ElementType = "container"
	ElementTypeEditor    ElementType = "editor"
	ElementTypeTree      ElementType = "tree"
)

type ElementTag string

const (
	ElementTagDiv    ElementTag = "div"
	ElementTagSpan   ElementTag = "span"
	ElementTagSvg    ElementTag = "svg"
	ElementTagImg    ElementTag = "img"
	ElementTagInput  ElementTag = "input"
	ElementTagIframe ElementTag = "iframe"
)

func Named(name string, ele *Element) *Element {
	if strings.HasSuffix(name, "Ele") {
		log.WarnLog("invalid name: %s", name)
	}
	ele.name = fmt.Sprintf("%sEle", name)
	return ele
}

func NewElement(type_ ElementType, tag ElementTag, children ...*Element) *Element {
	return &Element{
		type_: type_,
		tag:   tag,
		properties: map[string]string{
			"ch":      ".h - .borderTop - .borderBottom",
			"cw":      ".w - .borderLeft - .borderRight",
			"hovered": ".hoveredByMouse",
			"x2":      ".x + .w",
			"y2":      ".y + .h",
		},
		isStatic: map[string]bool{},
		children: children,
		name:     string(type_),
	}
}

type Element struct {
	type_      ElementType
	tag        ElementTag
	properties map[string]string
	isStatic   map[string]bool
	children   []*Element

	name          string
	selfIndex     []int
	isLocalRoot   bool
	localElements []*Element
}

func (e *Element) SelfIndex() []int {
	return e.selfIndex
}

func (e *Element) SetSelfIndex(index []int) {
	e.selfIndex = index
}

func (e *Element) Type() ElementType {
	return e.type_
}

func (e *Element) LocalElements() []*Element {
	return e.localElements
}

func (e *Element) AddLocalElement(ele *Element) {
	e.localElements = append(e.localElements, ele)
}

func (e *Element) IsLocalRoot() bool {
	return e.isLocalRoot
}

func (e *Element) SetLocalRoot() *Element {
	e.isLocalRoot = true
	return e
}

func (e *Element) Tag() ElementTag {
	return e.tag
}

func (e *Element) Name() string {
	return e.name
}

func (e *Element) Properties() map[string]string {
	return e.properties
}

func (e *Element) IsStatic(name string) bool {
	return e.isStatic[name]
}

func (e *Element) SetStaticProperty(k, v string) *Element {
	e.SetProperty(k, v)
	e.isStatic[k] = true
	return e
}

func (e *Element) SetProperty(k, v string) *Element {
	e.properties[k] = v
	return e
}

func (e *Element) Children() []*Element {
	return e.children
}
