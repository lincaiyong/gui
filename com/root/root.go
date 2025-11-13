package root

import (
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/js"
)

func Root(children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret, children...)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	_props map[string]string
}

func (c *Component) Code(code string) *Component {
	js.Set("Root", code)
	return c
}
