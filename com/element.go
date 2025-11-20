package com

import "strconv"

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
	ElementTagDiv    ElementTag  = "div"
	ElementTagSpan   ElementTag  = "span"
	ElementTagSvg    ElementTag  = "svg"
	ElementTagImg    ElementTag  = "img"
	ElementTagInput  ElementTag  = "input"
	ElementTagIframe ElementType = "iframe"
)

func NewElement(type_ ElementType, tag ElementTag, children ...*Element) *Element {
	return &Element{
		type_:         type_,
		tag:           tag,
		name:          string(tag),
		properties:    map[string]string{},
		methods:       map[string]string{},
		children:      children,
		localChildren: map[string][]int{},
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
	slots      []*Element

	localRoot     bool
	localChildren map[string][]int
}

func (e *Element) Type() ElementType {
	return e.type_
}

func (e *Element) SetType(type_ ElementType) {
	e.type_ = type_
}

func (e *Element) LocalChildren() map[string][]int {
	return e.localChildren
}

func (e *Element) SetLocalChildren(k string, v []int) {
	e.localChildren[k] = v
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

func (e *Element) Slots() []*Element {
	return e.slots
}

func (e *Element) Contains(slots ...*Element) *Element {
	e.slots = slots
	return e
}

func (e *Element) Position(s string) *Element {
	e.SetProperty("position", s)
	return e
}

func (e *Element) X(s string) *Element {
	e.SetProperty("x", s)
	return e
}

func (e *Element) Y(s string) *Element {
	e.SetProperty("y", s)
	return e
}

func (e *Element) W(s string) *Element {
	e.SetProperty("w", s)
	return e
}

func (e *Element) H(s string) *Element {
	e.SetProperty("h", s)
	return e
}

func (e *Element) V(s string) *Element {
	e.SetProperty("v", s)
	return e
}

func (e *Element) X2(s string) *Element {
	e.SetProperty("x2", s)
	return e
}

func (e *Element) Y2(s string) *Element {
	e.SetProperty("y2", s)
	return e
}

func (e *Element) Cw(s string) *Element {
	e.SetProperty("cw", s)
	return e
}

func (e *Element) Ch(s string) *Element {
	e.SetProperty("ch", s)
	return e
}

func (e *Element) BorderRadius(s string) *Element {
	e.SetProperty("borderRadius", s)
	return e
}

func (e *Element) Color(s string) *Element {
	e.SetProperty("color", s)
	return e
}

func (e *Element) BgColor(s string) *Element {
	e.SetProperty("backgroundColor", s)
	return e
}

func (e *Element) BorderColor(s string) *Element {
	e.SetProperty("borderColor", s)
	return e
}

func (e *Element) BoxShadow(s string) *Element {
	e.SetProperty("boxShadow", s)
	return e
}

func (e *Element) Background(s string) *Element {
	e.SetProperty("background", s)
	return e
}

func (e *Element) CaretColor(s string) *Element {
	e.SetProperty("caretColor", s)
	return e
}

func (e *Element) UserSelect(s string) *Element {
	e.SetProperty("userSelect", s)
	return e
}

func (e *Element) Cursor(s string) *Element {
	e.SetProperty("cursor", s)
	return e
}

func (e *Element) ZIndex(s string) *Element {
	e.SetProperty("zIndex", s)
	return e
}

func (e *Element) Opacity(s string) *Element {
	e.SetProperty("opacity", s)
	return e
}

func (e *Element) BorderStyle(s string) *Element {
	e.SetProperty("borderStyle", s)
	return e
}

func (e *Element) FontFamily(s string) *Element {
	e.SetProperty("fontFamily", s)
	return e
}

func (e *Element) FontSize(s string) *Element {
	e.SetProperty("fontSize", s)
	return e
}

func (e *Element) FontWeight(s string) *Element {
	e.SetProperty("fontWeight", s)
	return e
}

func (e *Element) Outline(s string) *Element {
	e.SetProperty("outline", s)
	return e
}

func (e *Element) LineHeight(s string) *Element {
	e.SetProperty("lineHeight", s)
	return e
}

func (e *Element) FontVariantLigatures(s string) *Element {
	e.SetProperty("fontVariantLigatures", s)
	return e
}

func (e *Element) InnerText(s string) *Element {
	e.SetProperty("innerText", s)
	return e
}

func (e *Element) ScrollTop(s string) *Element {
	e.SetProperty("scrollTop", s)
	return e
}

func (e *Element) ScrollLeft(s string) *Element {
	e.SetProperty("scrollLeft", s)
	return e
}

func (e *Element) BorderLeft(s int) *Element {
	e.SetProperty("borderLeft", strconv.Itoa(s))
	return e
}

func (e *Element) BorderRight(s int) *Element {
	e.SetProperty("borderRight", strconv.Itoa(s))
	return e
}

func (e *Element) BorderTop(s int) *Element {
	e.SetProperty("borderTop", strconv.Itoa(s))
	return e
}

func (e *Element) BorderBottom(s int) *Element {
	e.SetProperty("borderBottom", strconv.Itoa(s))
	return e
}

func (e *Element) Hovered(s string) *Element {
	e.SetProperty("hovered", s)
	return e
}

func (e *Element) HoveredByMouse(s string) *Element {
	e.SetProperty("hoveredByMouse", s)
	return e
}

func (e *Element) OnClick(s string) *Element {
	e.SetProperty("onClick", s)
	return e
}

func (e *Element) OnDoubleClick(s string) *Element {
	e.SetProperty("onDoubleClick", s)
	return e
}

func (e *Element) OnContextMenu(s string) *Element {
	e.SetProperty("onContextMenu", s)
	return e
}

func (e *Element) OnMouseDown(s string) *Element {
	e.SetProperty("onMouseDown", s)
	return e
}

func (e *Element) OnMouseMove(s string) *Element {
	e.SetProperty("onMouseMove", s)
	return e
}

func (e *Element) OnMouseUp(s string) *Element {
	e.SetProperty("onMouseUp", s)
	return e
}

func (e *Element) OnWheel(s string) *Element {
	e.SetProperty("onWheel", s)
	return e
}

func (e *Element) OnInput(s string) *Element {
	e.SetProperty("onInput", s)
	return e
}

func (e *Element) OnKeyUp(s string) *Element {
	e.SetProperty("onKeyUp", s)
	return e
}

func (e *Element) OnKeyDown(s string) *Element {
	e.SetProperty("onKeyDown", s)
	return e
}

func (e *Element) OnCompositionStart(s string) *Element {
	e.SetProperty("onCompositionStart", s)
	return e
}

func (e *Element) OnCompositionUpdate(s string) *Element {
	e.SetProperty("onCompositionUpdate", s)
	return e
}

func (e *Element) OnCompositionEnd(s string) *Element {
	e.SetProperty("onCompositionEnd", s)
	return e
}

func (e *Element) OnPaste(s string) *Element {
	e.SetProperty("onPaste", s)
	return e
}

func (e *Element) OnCopy(s string) *Element {
	e.SetProperty("onCopy", s)
	return e
}

func (e *Element) OnCut(s string) *Element {
	e.SetProperty("onCut", s)
	return e
}

func (e *Element) OnActive(s string) *Element {
	e.SetProperty("onActive", s)
	return e
}

func (e *Element) OnFocus(s string) *Element {
	e.SetProperty("onFocus", s)
	return e
}

func (e *Element) OnHover(s string) *Element {
	e.SetProperty("onHover", s)
	return e
}

func (e *Element) OnClickOutside(s string) *Element {
	e.SetProperty("onClickOutside", s)
	return e
}

func (e *Element) OnScrollTop(s string) *Element {
	e.SetProperty("onScrollTop", s)
	return e
}

func (e *Element) OnScrollLeft(s string) *Element {
	e.SetProperty("onScrollLeft", s)
	return e
}

func (e *Element) Placeholder(s string) *Element {
	e.SetProperty("placeholder", s)
	return e
}

func (e *Element) SrcDoc(s string) *Element {
	e.SetProperty("srcdoc", s)
	return e
}
