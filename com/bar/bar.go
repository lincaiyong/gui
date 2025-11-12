package bar

import (
	"github.com/lincaiyong/page/com"
)

func create() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.OnMouseDown("e.handleMouseDown").ZIndex("1")
	return ret
}

func VBar() *Component {
	ret := create()
	ret.Cursor("'col-resize'").W("20")
	return ret
}

func HBar() *Component {
	ret := create()
	ret.Cursor("'row-resize'").H("20")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	handleMouseDown com.Method
}
