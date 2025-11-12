package divider

import (
	"github.com/lincaiyong/page/com"
)

func VDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BgColor("'black'").W("1")
	return ret
}

func HDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BgColor("'black'").H("1")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
