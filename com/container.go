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
	ret.SetLocalRoot(true)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.SetSlots(child)
	ret.List(false).Virtual(false).Scrollable(false)
	ret.SetProperty("onUpdated", `function(k) {
    // items
    if (k === 'items' && this.list) {
        container_updateList.apply(this);
    }

    // scroll
    if (this.list && this.virtual && this.items instanceof Array) {
        if ((k === 'scrollLeft' || k === 'scrollTop') && this.items instanceof Array){
            container_updateList.apply(this);
        }
    } else if (this.list) {
        const RESERVED_COUNT = 2;
        if (k === 'scrollLeft') {
            for (let i = RESERVED_COUNT; i < this.children.length; i++) {
                const child = this.children[i];
                child.x = child.data.x - this.scrollLeft;
            }
        } else if (k === 'scrollTop') {
            for (let i = RESERVED_COUNT; i < this.children.length; i++) {
                const child = this.children[i];
                child.y = child.data.y - this.scrollTop;
            }
        }
    }

    // w & h -> 影响scroll
    if (this.scrollable) {
        if ((k === 'w' || k === 'h') && this.items instanceof Array) {
            container_updateList.apply(this);
        }
    }
}`)
	ret.SetMethod("onCreated", `function() {
    if (!this.list) {
        const child = g.createElement(this.model.slot[0], this);
        this.childWidth = child.w;
        this.childHeight = child.h;
        child.onUpdated((k, v) => {
            if (k === 'w') {
                this.childWidth = v;
            } else if (k === 'h') {
                this.childHeight = v;
            }
        });
    }

    if (this.scrollable) {
        this.hBar = new Scrollbar(this, 'h');
        this.vBar = new Scrollbar(this, 'v');
        const bars = [this.hBar, this.vBar];
        bars.forEach(bar => bar.initDraggable());
        this.onWheel = (_, ev) => {
            ev.preventDefault();
            bars.forEach(bar => bar.handleWheel(ev));
        };
    }
}`)
	return ret
}
