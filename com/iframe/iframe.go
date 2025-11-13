package iframe

import "github.com/lincaiyong/gui/com"

func Iframe() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("iframe", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	setHtml com.Method
}
