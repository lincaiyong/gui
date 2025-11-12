package img

import "github.com/lincaiyong/page/com"

func Img(src string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("img", ret)
	ret.Src(src)
	return ret
}

func Svg(src string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("svg", ret)
	ret.Src(src)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	src       com.Property `default:"''"`
	onUpdated com.Method
}

func (b *Component) Src(s string) *Component {
	b.SetProp("src", s)
	return b
}
