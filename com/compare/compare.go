package compare

import (
	"github.com/lincaiyong/page/com"
)

func Compare() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	onCreated com.Method
	_destroy  com.Method
}
