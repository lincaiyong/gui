package com

import "strconv"

func NewOpt() *Opt {
	ret := &Opt{}
	ret.BaseOpt = NewBaseOpt[Opt](ret)
	return ret
}

type Opt struct {
	*BaseOpt[Opt]
}

func NewBaseOpt[T any](self *T) *BaseOpt[T] {
	return &BaseOpt[T]{
		self:       self,
		properties: make(map[string]string),
	}
}

type BaseOpt[T any] struct {
	self       *T
	properties map[string]string
}

func (o *BaseOpt[T]) SetProperty(key, value string) *T {
	o.properties[key] = value
	return o.self
}

func (o *BaseOpt[T]) Properties() map[string]string {
	return o.properties
}

func (o *BaseOpt[T]) Init(e *Element) {
	if o == nil {
		return
	}
	for k, v := range o.properties {
		e.SetProperty(k, v)
	}
}

func (o *BaseOpt[T]) Position(s string) *T {
	o.SetProperty("position", s)
	return o.self
}

func (o *BaseOpt[T]) X(s string) *T {
	o.SetProperty("x", s)
	return o.self
}

func (o *BaseOpt[T]) Y(s string) *T {
	o.SetProperty("y", s)
	return o.self
}

func (o *BaseOpt[T]) W(s string) *T {
	o.SetProperty("w", s)
	return o.self
}

func (o *BaseOpt[T]) H(s string) *T {
	o.SetProperty("h", s)
	return o.self
}

func (o *BaseOpt[T]) V(s string) *T {
	o.SetProperty("v", s)
	return o.self
}

func (o *BaseOpt[T]) X2(s string) *T {
	o.SetProperty("x2", s)
	return o.self
}

func (o *BaseOpt[T]) Y2(s string) *T {
	o.SetProperty("y2", s)
	return o.self
}

func (o *BaseOpt[T]) Cw(s string) *T {
	o.SetProperty("cw", s)
	return o.self
}

func (o *BaseOpt[T]) Ch(s string) *T {
	o.SetProperty("ch", s)
	return o.self
}

func (o *BaseOpt[T]) BorderRadius(s string) *T {
	o.SetProperty("borderRadius", s)
	return o.self
}

func (o *BaseOpt[T]) Color(s string) *T {
	o.SetProperty("color", s)
	return o.self
}

func (o *BaseOpt[T]) BgColor(s string) *T {
	o.SetProperty("backgroundColor", s)
	return o.self
}

func (o *BaseOpt[T]) BorderColor(s string) *T {
	o.SetProperty("borderColor", s)
	return o.self
}

func (o *BaseOpt[T]) BoxShadow(s string) *T {
	o.SetProperty("boxShadow", s)
	return o.self
}

func (o *BaseOpt[T]) Background(s string) *T {
	o.SetProperty("background", s)
	return o.self
}

func (o *BaseOpt[T]) CaretColor(s string) *T {
	o.SetProperty("caretColor", s)
	return o.self
}

func (o *BaseOpt[T]) UserSelect(s string) *T {
	o.SetProperty("userSelect", s)
	return o.self
}

func (o *BaseOpt[T]) Cursor(s string) *T {
	o.SetProperty("cursor", s)
	return o.self
}

func (o *BaseOpt[T]) ZIndex(s string) *T {
	o.SetProperty("zIndex", s)
	return o.self
}

func (o *BaseOpt[T]) Opacity(s string) *T {
	o.SetProperty("opacity", s)
	return o.self
}

func (o *BaseOpt[T]) BorderStyle(s string) *T {
	o.SetProperty("borderStyle", s)
	return o.self
}

func (o *BaseOpt[T]) FontFamily(s string) *T {
	o.SetProperty("fontFamily", s)
	return o.self
}

func (o *BaseOpt[T]) FontSize(s string) *T {
	o.SetProperty("fontSize", s)
	return o.self
}

func (o *BaseOpt[T]) FontWeight(s string) *T {
	o.SetProperty("fontWeight", s)
	return o.self
}

func (o *BaseOpt[T]) Outline(s string) *T {
	o.SetProperty("outline", s)
	return o.self
}

func (o *BaseOpt[T]) LineHeight(s string) *T {
	o.SetProperty("lineHeight", s)
	return o.self
}

func (o *BaseOpt[T]) FontVariantLigatures(s string) *T {
	o.SetProperty("fontVariantLigatures", s)
	return o.self
}

func (o *BaseOpt[T]) InnerText(s string) *T {
	o.SetProperty("innerText", s)
	return o.self
}

func (o *BaseOpt[T]) ScrollTop(s string) *T {
	o.SetProperty("scrollTop", s)
	return o.self
}

func (o *BaseOpt[T]) ScrollLeft(s string) *T {
	o.SetProperty("scrollLeft", s)
	return o.self
}

func (o *BaseOpt[T]) BorderLeft(s int) *T {
	o.SetProperty("borderLeft", strconv.Itoa(s))
	return o.self
}

func (o *BaseOpt[T]) BorderRight(s int) *T {
	o.SetProperty("borderRight", strconv.Itoa(s))
	return o.self
}

func (o *BaseOpt[T]) BorderTop(s int) *T {
	o.SetProperty("borderTop", strconv.Itoa(s))
	return o.self
}

func (o *BaseOpt[T]) BorderBottom(s int) *T {
	o.SetProperty("borderBottom", strconv.Itoa(s))
	return o.self
}

func (o *BaseOpt[T]) Hovered(s string) *T {
	o.SetProperty("hovered", s)
	return o.self
}

func (o *BaseOpt[T]) HoveredByMouse(s string) *T {
	o.SetProperty("hoveredByMouse", s)
	return o.self
}

func (o *BaseOpt[T]) OnClick(s string) *T {
	o.SetProperty("onClick", s)
	return o.self
}

func (o *BaseOpt[T]) OnDoubleClick(s string) *T {
	o.SetProperty("onDoubleClick", s)
	return o.self
}

func (o *BaseOpt[T]) OnContextMenu(s string) *T {
	o.SetProperty("onContextMenu", s)
	return o.self
}

func (o *BaseOpt[T]) OnMouseDown(s string) *T {
	o.SetProperty("onMouseDown", s)
	return o.self
}

func (o *BaseOpt[T]) OnMouseMove(s string) *T {
	o.SetProperty("onMouseMove", s)
	return o.self
}

func (o *BaseOpt[T]) OnMouseUp(s string) *T {
	o.SetProperty("onMouseUp", s)
	return o.self
}

func (o *BaseOpt[T]) OnWheel(s string) *T {
	o.SetProperty("onWheel", s)
	return o.self
}

func (o *BaseOpt[T]) OnInput(s string) *T {
	o.SetProperty("onInput", s)
	return o.self
}

func (o *BaseOpt[T]) OnKeyUp(s string) *T {
	o.SetProperty("onKeyUp", s)
	return o.self
}

func (o *BaseOpt[T]) OnKeyDown(s string) *T {
	o.SetProperty("onKeyDown", s)
	return o.self
}

func (o *BaseOpt[T]) OnCompositionStart(s string) *T {
	o.SetProperty("onCompositionStart", s)
	return o.self
}

func (o *BaseOpt[T]) OnCompositionUpdate(s string) *T {
	o.SetProperty("onCompositionUpdate", s)
	return o.self
}

func (o *BaseOpt[T]) OnCompositionEnd(s string) *T {
	o.SetProperty("onCompositionEnd", s)
	return o.self
}

func (o *BaseOpt[T]) OnPaste(s string) *T {
	o.SetProperty("onPaste", s)
	return o.self
}

func (o *BaseOpt[T]) OnCopy(s string) *T {
	o.SetProperty("onCopy", s)
	return o.self
}

func (o *BaseOpt[T]) OnCut(s string) *T {
	o.SetProperty("onCut", s)
	return o.self
}

func (o *BaseOpt[T]) OnActive(s string) *T {
	o.SetProperty("onActive", s)
	return o.self
}

func (o *BaseOpt[T]) OnFocus(s string) *T {
	o.SetProperty("onFocus", s)
	return o.self
}

func (o *BaseOpt[T]) OnHover(s string) *T {
	o.SetProperty("onHover", s)
	return o.self
}

func (o *BaseOpt[T]) OnClickOutside(s string) *T {
	o.SetProperty("onClickOutside", s)
	return o.self
}

func (o *BaseOpt[T]) OnScrollTop(s string) *T {
	o.SetProperty("onScrollTop", s)
	return o.self
}

func (o *BaseOpt[T]) OnScrollLeft(s string) *T {
	o.SetProperty("onScrollLeft", s)
	return o.self
}

func (o *BaseOpt[T]) Placeholder(s string) *T {
	o.SetProperty("placeholder", s)
	return o.self
}

func (o *BaseOpt[T]) SrcDoc(s string) *T {
	o.SetProperty("srcdoc", s)
	return o.self
}

func (o *BaseOpt[T]) Src(s string) *T {
	o.SetProperty("src", s)
	return o.self
}
