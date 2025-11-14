package com

import "strconv"

func NewBaseComponent[T any](tag string, self *T, children ...Component) *BaseComponent[T] {
	return &BaseComponent[T]{
		extraInfo: ExtraInfo{
			bindJs:       make(map[string]string),
			defaultValue: make(map[string]string),
		},
		self:     self,
		tag:      tag,
		children: children,
		props: map[string]string{
			"ch":      ".h - .borderTop - .borderBottom",
			"cw":      ".w - .borderLeft - .borderRight",
			"hovered": ".hoveredByMouse",
			"x2":      ".x + .w",
			"y2":      ".y + .h",
		},
	}
}

type BaseComponent[T any] struct {
	extraInfo ExtraInfo
	self      *T
	name      string
	tag       string
	children  []Component
	slots     []Component
	props     map[string]string
}

func (b *BaseComponent[T]) Name() string {
	return b.name
}

func (b *BaseComponent[T]) ExtraInfo() *ExtraInfo {
	return &b.extraInfo
}

func (b *BaseComponent[T]) Tag() string {
	return b.tag
}

func (b *BaseComponent[T]) Children() []Component {
	return b.children
}

func (b *BaseComponent[T]) Slots() []Component {
	return b.slots
}

func (b *BaseComponent[T]) Props() map[string]string {
	return b.props
}

func (b *BaseComponent[T]) SetProp(k, v string) {
	b.props[k] = v
}

func (b *BaseComponent[T]) NameAs(name string) *T {
	b.name = name
	return b.self
}

func (b *BaseComponent[T]) SetSlots(s ...Component) *T {
	b.slots = append(b.slots, s...)
	return b.self
}

func (b *BaseComponent[T]) OnCreated(fn string) *T {
	b.SetProp("onCreatedFn", fn)
	return b.self
}

func (b *BaseComponent[T]) OnUpdated(fn string) *T {
	b.SetProp("onUpdatedFn", fn)
	return b.self
}

func (b *BaseComponent[T]) Position(s string) *T {
	b.SetProp("position", s)
	return b.self
}

func (b *BaseComponent[T]) X(s string) *T {
	b.SetProp("x", s)
	return b.self
}

func (b *BaseComponent[T]) Y(s string) *T {
	b.SetProp("y", s)
	return b.self
}

func (b *BaseComponent[T]) W(s string) *T {
	b.SetProp("w", s)
	return b.self
}

func (b *BaseComponent[T]) H(s string) *T {
	b.SetProp("h", s)
	return b.self
}

func (b *BaseComponent[T]) V(s string) *T {
	b.SetProp("v", s)
	return b.self
}

func (b *BaseComponent[T]) X2(s string) *T {
	b.SetProp("x2", s)
	return b.self
}

func (b *BaseComponent[T]) Y2(s string) *T {
	b.SetProp("y2", s)
	return b.self
}

func (b *BaseComponent[T]) Cw(s string) *T {
	b.SetProp("cw", s)
	return b.self
}

func (b *BaseComponent[T]) Ch(s string) *T {
	b.SetProp("ch", s)
	return b.self
}

func (b *BaseComponent[T]) BorderRadius(s string) *T {
	b.SetProp("borderRadius", s)
	return b.self
}

func (b *BaseComponent[T]) Color(s string) *T {
	b.SetProp("color", s)
	return b.self
}

func (b *BaseComponent[T]) BgColor(s string) *T {
	b.SetProp("backgroundColor", s)
	return b.self
}

func (b *BaseComponent[T]) BorderColor(s string) *T {
	b.SetProp("borderColor", s)
	return b.self
}

func (b *BaseComponent[T]) BoxShadow(s string) *T {
	b.SetProp("boxShadow", s)
	return b.self
}

func (b *BaseComponent[T]) Background(s string) *T {
	b.SetProp("background", s)
	return b.self
}

func (b *BaseComponent[T]) CaretColor(s string) *T {
	b.SetProp("caretColor", s)
	return b.self
}

func (b *BaseComponent[T]) UserSelect(s string) *T {
	b.SetProp("userSelect", s)
	return b.self
}

func (b *BaseComponent[T]) Cursor(s string) *T {
	b.SetProp("cursor", s)
	return b.self
}

func (b *BaseComponent[T]) ZIndex(s string) *T {
	b.SetProp("zIndex", s)
	return b.self
}

func (b *BaseComponent[T]) Opacity(s string) *T {
	b.SetProp("opacity", s)
	return b.self
}

func (b *BaseComponent[T]) BorderStyle(s string) *T {
	b.SetProp("borderStyle", s)
	return b.self
}

func (b *BaseComponent[T]) FontFamily(s string) *T {
	b.SetProp("fontFamily", s)
	return b.self
}

func (b *BaseComponent[T]) FontSize(s string) *T {
	b.SetProp("fontSize", s)
	return b.self
}

func (b *BaseComponent[T]) FontWeight(s string) *T {
	b.SetProp("fontWeight", s)
	return b.self
}

func (b *BaseComponent[T]) Outline(s string) *T {
	b.SetProp("outline", s)
	return b.self
}

func (b *BaseComponent[T]) LineHeight(s string) *T {
	b.SetProp("lineHeight", s)
	return b.self
}

func (b *BaseComponent[T]) FontVariantLigatures(s string) *T {
	b.SetProp("fontVariantLigatures", s)
	return b.self
}

func (b *BaseComponent[T]) InnerText(s string) *T {
	b.SetProp("innerText", s)
	return b.self
}

func (b *BaseComponent[T]) ScrollTop(s string) *T {
	b.SetProp("scrollTop", s)
	return b.self
}

func (b *BaseComponent[T]) ScrollLeft(s string) *T {
	b.SetProp("scrollLeft", s)
	return b.self
}

func (b *BaseComponent[T]) BorderLeft(s int) *T {
	b.SetProp("borderLeft", strconv.Itoa(s))
	return b.self
}

func (b *BaseComponent[T]) BorderRight(s int) *T {
	b.SetProp("borderRight", strconv.Itoa(s))
	return b.self
}

func (b *BaseComponent[T]) BorderTop(s int) *T {
	b.SetProp("borderTop", strconv.Itoa(s))
	return b.self
}

func (b *BaseComponent[T]) BorderBottom(s int) *T {
	b.SetProp("borderBottom", strconv.Itoa(s))
	return b.self
}

func (b *BaseComponent[T]) Hovered(s string) *T {
	b.SetProp("hovered", s)
	return b.self
}

func (b *BaseComponent[T]) HoveredByMouse(s string) *T {
	b.SetProp("hoveredByMouse", s)
	return b.self
}

func (b *BaseComponent[T]) OnClick(s string) *T {
	b.SetProp("onClick", s)
	return b.self
}

func (b *BaseComponent[T]) OnDoubleClick(s string) *T {
	b.SetProp("onDoubleClick", s)
	return b.self
}

func (b *BaseComponent[T]) OnContextMenu(s string) *T {
	b.SetProp("onContextMenu", s)
	return b.self
}

func (b *BaseComponent[T]) OnMouseDown(s string) *T {
	b.SetProp("onMouseDown", s)
	return b.self
}

func (b *BaseComponent[T]) OnMouseMove(s string) *T {
	b.SetProp("onMouseMove", s)
	return b.self
}

func (b *BaseComponent[T]) OnMouseUp(s string) *T {
	b.SetProp("onMouseUp", s)
	return b.self
}

func (b *BaseComponent[T]) OnWheel(s string) *T {
	b.SetProp("onWheel", s)
	return b.self
}

func (b *BaseComponent[T]) OnInput(s string) *T {
	b.SetProp("onInput", s)
	return b.self
}

func (b *BaseComponent[T]) OnKeyUp(s string) *T {
	b.SetProp("onKeyUp", s)
	return b.self
}

func (b *BaseComponent[T]) OnKeyDown(s string) *T {
	b.SetProp("onKeyDown", s)
	return b.self
}

func (b *BaseComponent[T]) OnCompositionStart(s string) *T {
	b.SetProp("onCompositionStart", s)
	return b.self
}

func (b *BaseComponent[T]) OnCompositionUpdate(s string) *T {
	b.SetProp("onCompositionUpdate", s)
	return b.self
}

func (b *BaseComponent[T]) OnCompositionEnd(s string) *T {
	b.SetProp("onCompositionEnd", s)
	return b.self
}

func (b *BaseComponent[T]) OnPaste(s string) *T {
	b.SetProp("onPaste", s)
	return b.self
}

func (b *BaseComponent[T]) OnCopy(s string) *T {
	b.SetProp("onCopy", s)
	return b.self
}

func (b *BaseComponent[T]) OnCut(s string) *T {
	b.SetProp("onCut", s)
	return b.self
}

func (b *BaseComponent[T]) OnActive(s string) *T {
	b.SetProp("onActive", s)
	return b.self
}

func (b *BaseComponent[T]) OnFocus(s string) *T {
	b.SetProp("onFocus", s)
	return b.self
}

func (b *BaseComponent[T]) OnHover(s string) *T {
	b.SetProp("onHover", s)
	return b.self
}

func (b *BaseComponent[T]) OnClickOutside(s string) *T {
	b.SetProp("onClickOutside", s)
	return b.self
}

func (b *BaseComponent[T]) OnScrollTop(s string) *T {
	b.SetProp("onScrollTop", s)
	return b.self
}

func (b *BaseComponent[T]) OnScrollLeft(s string) *T {
	b.SetProp("onScrollLeft", s)
	return b.self
}
