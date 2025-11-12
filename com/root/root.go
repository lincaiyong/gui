package root

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
)

func Root(children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret, children...)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}

func (c *Component) Code(code string) *Component {
	js.Set("Root", code)
	return c
}
