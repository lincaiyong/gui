package com

func Div() *Element {
	return NewElement(ElementTypeDiv, ElementTagDiv)
}

func Text(text string) *Element {
	ret := NewElement(ElementTypeText, ElementTagSpan)
	ret.InnerText(text).FontSize("Math.floor(.h * 2 / 3)").LineHeight(".h").
		W("g.textWidth(.innerText, .fontFamily, .fontSize)")
	return ret
}

func Input() *Element {
	ret := NewElement(ElementTypeInput, ElementTagInput)
	ret.LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)")
	return ret
}

func Img(src string) *Element {
	ret := NewElement(ElementTypeImg, ElementTagImg)
	ret.SetProperty("src", src)
	return ret
}

func Svg(src string) *Element {
	ret := Img(src)
	ret.SetType(ElementTypeSvg)
	ret.SetTag(ElementTagSvg)
	return ret
}

func VDivider() *Element {
	ret := NewElement(ElementTypeDivider, ElementTagDiv)
	ret.BgColor("'black'").W("1")
	return ret
}

func HDivider() *Element {
	ret := NewElement(ElementTypeDivider, ElementTagDiv)
	ret.BgColor("'black'").H("1")
	return ret
}

func VBar() *Element {
	ret := NewElement(ElementTypeBar, ElementTagDiv)
	ret.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'col-resize'").W("20")
	return ret
}

func HBar() *Element {
	ret := NewElement(ElementTypeBar, ElementTagDiv)
	ret.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'row-resize'").H("20")
	return ret
}

func VScrollbar() *Element {
	ret := NewElement(ElementTypeScrollbar, ElementTagDiv)
	ret.ZIndex("1").
		BgColor("'#7f7e80'").
		Opacity("0.5").
		BorderRadius(".w / 2").
		Cursor("'default'").
		X(".vertical ? parent.cw - parent.scrollBarMargin - parent.scrollBarWidth : 0").
		Y(".vertical ? 0 : parent.ch - parent.scrollBarMargin - parent.scrollBarWidth").
		W(".vertical ? parent.scrollBarWidth : 0").
		H(".vertical ? 0 : parent.scrollBarWidth").
		V("0").
		SetProperty("vertical", "true")
	return ret
}

func HScrollbar() *Element {
	ret := HScrollbar().SetProperty("vertical", "false")
	return ret
}

func Iframe() *Element {
	return NewElement(ElementTypeIframe, ElementTagIframe)
}
