package com

import "strconv"

func NewElement(type_, tag string, children ...*Element) *Element {
	return &Element{
		type_:      type_,
		tag:        tag,
		name:       tag,
		properties: map[string]string{},
		methods:    map[string]string{},
		children:   children,
	}
}

type Element struct {
	type_      string
	tag        string
	name       string
	depth      int
	properties map[string]string
	methods    map[string]string
	children   []*Element
	slots      []*Element
}

func (e *Element) SetTag(tag string) {
	e.tag = tag
}

func (e *Element) Tag() string {
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
