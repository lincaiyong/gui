package com

func Div() *Element {
	return NewElement("div", "div")
}

func Text(text string) *Element {
	ret := NewElement("text", "span")
	ret.InnerText(text).FontSize("Math.floor(.h * 2 / 3)").LineHeight(".h").
		W("g.textWidth(.innerText, .fontFamily, .fontSize)")
	return ret
}

func Input() *Element {
	ret := NewElement("input", "input")
	ret.LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)")
	return ret
}

func Img(src string) *Element {
	ret := NewElement("img", "img")
	ret.SetProperty("src", src)
	ret.SetMethod("onUpdated", `function(k, v) {
    if (k === 'src' && this.tag === 'svg') {
		g.fetchRes('./res/' + v).then(data => this.ref.innerHTML = data).catch(err => g.log.error(err));
	} else if (k === 'src' && this.tag === 'img') {
		this.ref.setAttribute(k, './res/' + v);
	}
}`)
	return ret
}

func Svg(src string) *Element {
	ret := Img(src)
	ret.SetTag("svg")
	return ret
}

func VDivider() *Element {
	ret := NewElement("divider", "div")
	ret.BgColor("'black'").W("1")
	return ret
}

func HDivider() *Element {
	ret := NewElement("divider", "div")
	ret.BgColor("'black'").H("1")
	return ret
}

func VBar() *Element {
	ret := NewElement("bar", "div")
	ret.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'col-resize'").W("20")
	return ret
}

func HBar() *Element {
	ret := NewElement("bar", "div")
	ret.OnMouseDown("bar_handleMouseDown").ZIndex("1").Cursor("'row-resize'").H("20")
	return ret
}

func VScrollbar() *Element {
	ret := NewElement("scrollbar", "div")
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
