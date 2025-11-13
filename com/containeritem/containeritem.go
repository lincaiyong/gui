package containeritem

import (
	"github.com/lincaiyong/gui/com"
)

func ContainerItem(children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret, children...)
	ret.Y("0").X("0").H("20")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	data      com.Property `default:"undefined"`
	compute   com.Property `default:"null"`
	update    com.Property `default:"null"`
	click     com.Property `default:"null"`
	onUpdated com.Method
}

func (b *Component) Compute(s string) *Component {
	b.SetProp("compute", s)
	return b
}

func (b *Component) Update(s string) *Component {
	b.SetProp("update", s)
	return b
}
