package input

import "github.com/lincaiyong/page/com"

func Input() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("input", ret)
	ret.LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	placeholder com.Property `default:"''"`
	onUpdated   com.Method
}

func (b *Component) Placeholder(s string) *Component {
	b.SetProp("placeholder", s)
	return b
}
