package div

import "github.com/lincaiyong/gui/com"

func Div() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
