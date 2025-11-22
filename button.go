package gui

func NewButtonOpt() *ButtonOpt {
	ret := &ButtonOpt{}
	ret.BaseOpt = NewBaseOpt[ButtonOpt](ret)
	ret.Svg(SvgProject).
		Selected("false").
		Text("''").
		Flag("false")
	return ret
}

type ButtonOpt struct {
	*BaseOpt[ButtonOpt]
}

func (o *ButtonOpt) Svg(s string) *ButtonOpt      { o.SetProperty("svg", s); return o }
func (o *ButtonOpt) Selected(s string) *ButtonOpt { o.SetProperty("selected", s); return o }
func (o *ButtonOpt) Text(s string) *ButtonOpt     { o.SetProperty("text", s); return o }
func (o *ButtonOpt) Flag(s string) *ButtonOpt     { o.SetProperty("flag", s); return o }

func Button(opt *ButtonOpt) *Element {
	ret := NewElement(ElementTypeButton, ElementTagDiv,
		Svg(NewOpt().X("4").Y(".x").W("parent.w - 2 * .x").H(".w").Color("parent.color"), "parent.svg"),
		Div(NewOpt().X("prev.x2 - .w + 1").Y("prev.y - 1").W("6").H(".w").V("0").BorderRadius("3")),
	)
	ret.SetLocalRoot()
	opt.W("24").H(".w").BorderRadius("6").
		BgColor(".selected ? '#3475F0' : ''").
		Color(".selected ? '#FFFFFF' : '6C707E'").
		OnHover("button_handleHover").
		OnActive("button_handleActive").
		Init(ret)
	return ret
}

func SourceRootButton(opt *ButtonOpt) *Element {
	ret := NewElement(ElementTypeButton, ElementTagDiv,
		Text(NewOpt().X("next.x2+4").Y("1").H("parent.h-2").Color(ColorGray110), "parent.text"),
		Svg(NewOpt().X(".y-2").Y("parent.h/2-.h/2+1").W("8").H(".w").Color("parent.color"), SvgSourceRootFileLayer),
	)
	ret.SetLocalRoot()
	opt.W("child.w + 21").H("20").BorderRadius("3").
		BgColor("''").Color("'6C707E'").
		OnHover("button_handleHover").
		OnActive("button_handleActive").
		Init(ret)
	return ret
}

func SourceDirButton(opt *ButtonOpt) *Element {
	ret := NewElement(ElementTypeButton, ElementTagDiv,
		Text(NewOpt().X("2").Y("1").H("parent.h-2").Color(ColorGray110), "parent.text"),
	)
	ret.SetLocalRoot()
	opt.W("child.w + 4").H("20").BorderRadius("3").
		BgColor("''").Color("'6C707E'").
		OnHover("button_handleHover").
		OnActive("button_handleActive").
		Init(ret)
	return ret
}

func SourceFileButton(opt *ButtonOpt) *Element {
	ret := NewElement(ElementTypeButton, ElementTagDiv,
		Text(NewOpt().X("next.x2+4").Y("1").H("parent.h-2").Color(ColorGray110), "parent.text"),
		Svg(NewOpt().X("3").Y("2").W("16").H(".w"), SvgGo),
	)
	ret.SetLocalRoot()
	opt.W("child.w + 26").H("20").BorderRadius("3").
		BgColor("''").Color("'6C707E'").
		OnHover("button_handleHover").
		OnActive("button_handleActive").
		Init(ret)
	return ret
}

func ToolButton(opt *ButtonOpt) *Element {
	ret := NewElement(ElementTypeButton, ElementTagDiv,
		Svg(NewOpt().X("4").Y(".x").W("parent.w - 2 * .x").H(".w").Color("parent.color"), "parent.svg"),
		Div(NewOpt().X("prev.x2 - .w + 1").Y("prev.y - 1").W("8").H(".w").V("parent.flag").BorderRadius("4").BgColor(ColorOrange).
			BorderColor(ColorWhite).BorderLeft(1).BorderRight(1).BorderTop(1).BorderBottom(1)),
	)
	ret.SetLocalRoot()
	opt.W("24").H(".w").BorderRadius("6").
		BgColor("''").Color("'6C707E'").
		OnHover("button_handleHover").
		OnActive("button_handleActive").
		Init(ret)
	return ret
}
