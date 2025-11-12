package divider

import (
	"github.com/lincaiyong/page/com"
)

func VDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BgColor("'black'").W("2")
	return ret
}

func HDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BgColor("'black'").H("2")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
