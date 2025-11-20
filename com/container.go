package com

import (
	"fmt"
	"strconv"
)

func NewContainerOpt() ContainerOpt {
	ret := ContainerOpt{
		BaseOpt: NewBaseOpt(),
	}
	ret.SetProperty("align", "'none'")
	ret.SetProperty("childHeight", "0")
	ret.SetProperty("childWidth", "0")
	ret.SetProperty("items", "[]")
	ret.SetProperty("list", "false")
	ret.SetProperty("minWidth", "0")
	ret.SetProperty("reuseItem", "false")
	ret.SetProperty("scrollBarFadeTime", "500")
	ret.SetProperty("scrollBarMinLength", "20")
	ret.SetProperty("scrollBarWidth", "6")
	ret.SetProperty("scrollBarMargin", "0")
	ret.SetProperty("scrollable", "true")
	ret.SetProperty("virtual", "false")
	return ret
}

type ContainerOpt struct {
	*BaseOpt
}

func (o *ContainerOpt) Align(s string) *ContainerOpt       { o.SetProperty("align", s); return o }
func (o *ContainerOpt) ChildHeight(s string) *ContainerOpt { o.SetProperty("childHeight", s); return o }
func (o *ContainerOpt) ChildWidth(s string) *ContainerOpt  { o.SetProperty("childWidth", s); return o }
func (o *ContainerOpt) Items(s string) *ContainerOpt       { o.SetProperty("items", s); return o }
func (o *ContainerOpt) List(s string) *ContainerOpt        { o.SetProperty("list", s); return o }
func (o *ContainerOpt) MinWidth(s string) *ContainerOpt    { o.SetProperty("minWidth", s); return o }
func (o *ContainerOpt) ReuseItem(s string) *ContainerOpt   { o.SetProperty("reuseItem", s); return o }
func (o *ContainerOpt) ScrollBarFadeTime(s string) *ContainerOpt {
	o.SetProperty("scrollBarFadeTime", s)
	return o
}
func (o *ContainerOpt) ScrollBarMinLength(s string) *ContainerOpt {
	o.SetProperty("scrollBarMinLength", s)
	return o
}
func (o *ContainerOpt) ScrollBarWidth(s string) *ContainerOpt {
	o.SetProperty("scrollBarWidth", s)
	return o
}
func (o *ContainerOpt) ScrollBarMargin(s string) *ContainerOpt {
	o.SetProperty("scrollBarMargin", s)
	return o
}
func (o *ContainerOpt) Scrollable(s string) *ContainerOpt { o.SetProperty("scrollable", s); return o }
func (o *ContainerOpt) Virtual(s string) *ContainerOpt    { o.SetProperty("virtual", s); return o }

//onCreated          Method
//_updateList        Method
//onUpdated          Method
//
//itemComp *containeritem.Element

func VListContainer(children ...Element) *Element {
	ret := &Element{}
	ret.BaseComponent = NewBaseComponent[Element]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.itemComp = containeritem.ContainerItem(children...)
	ret.SetSlots(ret.itemComp)
	ret.List(true).Virtual(true).Scrollable(true)
	return ret
}

func ListContainer(children ...Element) *Element {
	ret := &Element{}
	ret.BaseComponent = NewBaseComponent[Element]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.itemComp = containeritem.ContainerItem(children...)
	ret.SetSlots(ret.itemComp)
	ret.List(true).Virtual(false).Scrollable(true)
	return ret
}

func Container(child Element) *Element {
	ret := NewElement("container", "div"
		HScrollbar().NameAs("hBarEle"),
		VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.SetSlots(child)
	ret.List(false).Virtual(false).Scrollable(false)
	return ret
}
