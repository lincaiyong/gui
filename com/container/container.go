package container

import (
	"fmt"
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/com/containeritem"
	"github.com/lincaiyong/gui/com/scrollbar"
	"strconv"
)

func VListContainer(children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.itemComp = containeritem.ContainerItem(children...)
	ret.SetSlots(ret.itemComp)
	ret.List(true).Virtual(true).Scrollable(true)
	return ret
}

func ListContainer(children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.itemComp = containeritem.ContainerItem(children...)
	ret.SetSlots(ret.itemComp)
	ret.List(true).Virtual(false).Scrollable(true)
	return ret
}

func Container(child com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.SetSlots(child)
	ret.List(false).Virtual(false).Scrollable(false)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	align              com.Property `default:"'none'"`
	childHeight        com.Property `default:"0"`
	childWidth         com.Property `default:"0"`
	items              com.Property `default:"[]"`
	list               com.Property `default:"false"`
	minWidth           com.Property `default:"0"`
	reuseItem          com.Property `default:"false"`
	scrollBarFadeTime  com.Property `default:"500"`
	scrollBarMinLength com.Property `default:"20"`
	scrollBarWidth     com.Property `default:"6"`
	scrollBarMargin    com.Property `default:"0"`
	scrollable         com.Property `default:"true"`
	virtual            com.Property `default:"false"`
	onCreated          com.Method
	_updateList        com.Method
	onUpdated          com.Method

	itemComp *containeritem.Component
}

func (b *Component) ItemCompute(s string) *Component {
	if b.itemComp != nil {
		b.itemComp.Compute(s)
	}
	return b
}

func (b *Component) ItemOnUpdated(s string) *Component {
	if b.itemComp != nil {
		b.itemComp.Update(s)
	}
	return b
}

func (b *Component) ItemOnClick(s string) *Component {
	if b.itemComp != nil {
		b.itemComp.OnClick(s)
	}
	return b
}

func (b *Component) Align(s string) *Component {
	b.SetProp("align", s)
	return b
}

func (b *Component) ChildHeight(v int) *Component {
	b.SetProp("childHeight", strconv.Itoa(v))
	return b
}

func (b *Component) ChildWidth(v int) *Component {
	b.SetProp("childWidth", strconv.Itoa(v))
	return b
}

func (b *Component) List(v bool) *Component {
	b.SetProp("list", fmt.Sprintf("%v", v))
	return b
}

func (b *Component) MinWidth(v int) *Component {
	b.SetProp("minWidth", strconv.Itoa(v))
	return b
}

func (b *Component) ReuseItem(v bool) *Component {
	b.SetProp("reuseItem", fmt.Sprintf("%v", v))
	return b
}

func (b *Component) ScrollBarFadeTime(v int) *Component {
	b.SetProp("scrollBarFadeTime", strconv.Itoa(v))
	return b
}

func (b *Component) ScrollBarMinLength(v int) *Component {
	b.SetProp("scrollBarMinLength", strconv.Itoa(v))
	return b
}

func (b *Component) ScrollBarWidth(v int) *Component {
	b.SetProp("scrollBarWidth", strconv.Itoa(v))
	return b
}

func (b *Component) Scrollable(v bool) *Component {
	b.SetProp("scrollable", fmt.Sprintf("%v", v))
	return b
}

func (b *Component) Virtual(v bool) *Component {
	b.SetProp("virtual", fmt.Sprintf("%v", v))
	return b
}
