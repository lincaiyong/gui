package text

import "github.com/lincaiyong/gui/com"

func Text(text string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("span", ret)
	ret.InnerText(text)
	ret.FontSize("Math.floor(.h * 2 / 3)").
		LineHeight(".h").
		W("g.util.textWidth(.innerText, .fontFamily, .fontSize)")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	//align     com.Property `default:"'left'"`
	//onUpdated com.Method
}

//
//func (b *Component) Align(s string) *Component {
//	b.SetProp("align", s)
//	return b
//}
