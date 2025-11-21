package gui

func Div(opt *Opt, children ...*Element) *Element {
	ret := NewElement(ElementTypeDiv, ElementTagDiv, children...)
	opt.Init(ret)
	return ret
}

func Text(opt *Opt, text string) *Element {
	ret := NewElement(ElementTypeText, ElementTagSpan)
	opt.InnerText(text).FontSize("Math.floor(.h * 2 / 3)").LineHeight(".h").
		W("g.textWidth(.innerText, .fontFamily, .fontSize)").Init(ret)
	return ret
}

func Input(opt *Opt, placeholder string) *Element {
	ret := NewElement(ElementTypeInput, ElementTagInput)
	opt.Placeholder(placeholder).LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)").Init(ret)
	return ret
}

func Img(opt *Opt, src string) *Element {
	ret := NewElement(ElementTypeImg, ElementTagImg)
	opt.Src(src).Init(ret)
	return ret
}

func Svg(opt *Opt, src string) *Element {
	ret := NewElement(ElementTypeSvg, ElementTagSvg)
	opt.Src(src).Init(ret)
	return ret
}

func VDivider(opt *Opt) *Element {
	ret := NewElement(ElementTypeDivider, ElementTagDiv)
	opt.BgColor("'black'").W("1").Init(ret)
	return ret
}

func HDivider(opt *Opt) *Element {
	ret := NewElement(ElementTypeDivider, ElementTagDiv)
	opt.BgColor("'black'").H("1").Init(ret)
	return ret
}

func VBar(opt *Opt) *Element {
	ret := NewElement(ElementTypeBar, ElementTagDiv)
	opt.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'col-resize'").W("20").Init(ret)
	return ret
}

func HBar(opt *Opt) *Element {
	ret := NewElement(ElementTypeBar, ElementTagDiv)
	opt.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'row-resize'").H("20").Init(ret)
	return ret
}

func VScrollbar(opt *Opt) *Element {
	ret := NewElement(ElementTypeScrollbar, ElementTagDiv)
	opt.ZIndex("1").
		BgColor("'#7f7e80'").
		Opacity("0.5").
		BorderRadius(".w / 2").
		Cursor("'default'").
		X(".vertical ? parent.cw - parent.scrollBarMargin - parent.scrollBarWidth : 0").
		Y(".vertical ? 0 : parent.ch - parent.scrollBarMargin - parent.scrollBarWidth").
		W(".vertical ? parent.scrollBarWidth : 0").
		H(".vertical ? 0 : parent.scrollBarWidth").
		V("0").
		SetProperty("vertical", "true").
		Init(ret)
	return ret
}

func HScrollbar(opt *Opt) *Element {
	ret := VScrollbar(opt)
	opt.SetProperty("vertical", "false").Init(ret)
	return ret
}

func Iframe(opt *Opt) *Element {
	ret := NewElement(ElementTypeIframe, ElementTagIframe)
	opt.Init(ret)
	return ret
}
