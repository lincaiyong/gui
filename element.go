package gui

import (
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

func NewElement(type_ ElementType, tag ElementTag, children ...*Element) *Element {
	return &Element{
		type_:      type_,
		tag:        tag,
		properties: map[string]string{},
		methods:    map[string]string{},
		children:   children,
		name:       string(tag),
	}
}

type Element struct {
	type_      ElementType
	tag        ElementTag
	properties map[string]string
	methods    map[string]string
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

func (e *Element) SetType(type_ ElementType) {
	e.type_ = type_
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

func (e *Element) SetLocalRoot() {
	e.isLocalRoot = true
}

func (e *Element) SetTag(tag ElementTag) {
	e.tag = tag
}

func (e *Element) Tag() ElementTag {
	return e.tag
}

func (e *Element) Name() string {
	return e.name
}

func (e *Element) SetName(name string) *Element {
	if !strings.HasSuffix(name, "Ele") {
		log.FatalLog("local name must end with 'Ele': %s", name)
	}
	e.name = name
	return e
}

func (e *Element) Properties() map[string]string {
	return e.properties
}

func (e *Element) SetProperty(k, v string) *Element {
	e.properties[k] = v
	return e
}

func (e *Element) Methods() map[string]string {
	return e.methods
}

func (e *Element) SetMethod(k, v string) *Element {
	e.methods[k] = v
	return e
}

func (e *Element) Children() []*Element {
	return e.children
}
