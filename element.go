package gui

import (
	"github.com/lincaiyong/log"
	"strings"
)

type ElementType string

const (
	ElementTypeDiv           ElementType = "div"
	ElementTypeText          ElementType = "text"
	ElementTypeInput         ElementType = "input"
	ElementTypeImg           ElementType = "img"
	ElementTypeSvg           ElementType = "svg"
	ElementTypeDivider       ElementType = "divider"
	ElementTypeBar           ElementType = "bar"
	ElementTypeScrollbar     ElementType = "scrollbar"
	ElementTypeIframe        ElementType = "iframe"
	ElementTypeButton        ElementType = "button"
	ElementTypeCompare       ElementType = "compare"
	ElementTypeContainer     ElementType = "container"
	ElementTypeContainerItem ElementType = "container_item"
	ElementTypeEditor        ElementType = "editor"
	ElementTypeTree          ElementType = "tree"
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
		name:       string(tag),
		properties: map[string]string{},
		methods:    map[string]string{},
		children:   children,
	}
}

type Element struct {
	type_      ElementType
	tag        ElementTag
	name       string
	depth      int
	properties map[string]string
	methods    map[string]string
	children   []*Element
	slot       *Element

	localRoot     bool
	localChildren []*Element
	localName     string
	localIndex    []int
}

func (e *Element) LocalIndex() []int {
	return e.localIndex
}

func (e *Element) SetLocalIndex(localIndex []int) {
	e.localIndex = localIndex
}

func (e *Element) LocalName() string {
	return e.localName
}

func (e *Element) SetLocalName(localName string) *Element {
	if !strings.HasSuffix(localName, "Ele") {
		log.FatalLog("local name must end with 'Ele': %s", localName)
	}
	e.localName = localName
	return e
}

func (e *Element) SetDepth(depth int) {
	e.depth = depth
}

func (e *Element) Type() ElementType {
	return e.type_
}

func (e *Element) SetType(type_ ElementType) {
	e.type_ = type_
}

func (e *Element) LocalChildren() []*Element {
	return e.localChildren
}

func (e *Element) AddLocalChildren(ele *Element) {
	e.localChildren = append(e.localChildren, ele)
}

func (e *Element) LocalRoot() bool {
	return e.localRoot
}

func (e *Element) SetLocalRoot(localRoot bool) {
	e.localRoot = localRoot
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
	e.name = name
	return e
}

func (e *Element) Depth() int {
	return e.depth
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

func (e *Element) Slot() *Element {
	return e.slot
}

func (e *Element) SetSlot(slot *Element) *Element {
	e.slot = slot
	return e
}
