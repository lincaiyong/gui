package toyeditor

import (
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/com/container"
	"github.com/lincaiyong/gui/com/div"
	"github.com/lincaiyong/gui/com/input"
)

func ToyEditor() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		div.Div(),
		container.VListContainer(),
		div.Div().SetSlots(
			container.VListContainer(),
			container.VListContainer(),
			div.Div().SetSlots(
				input.Input(),
			),
			container.VListContainer(),
		),
	)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
