package text

import "github.com/lincaiyong/page/com"

func Text(text string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("span", ret)
	ret.SetProp("text", text)
	ret.FontSize("Math.floor(.h * 2 / 3)").
		LineHeight(".h").
		W("page.util.textWidth(.text, .fontFamily, .fontSize)")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	align     com.Property `default:"'left'"`
	text      com.Property `default:"''"`
	onUpdated com.Method
}

func (b *Component) Align(s string) *Component {
	b.SetProp("align", s)
	return b
}
