package gui

func NewContainerOpt() *ContainerOpt {
	ret := &ContainerOpt{}
	ret.BaseOpt = NewBaseOpt[ContainerOpt](ret)
	ret.Align("'none'").
		ChildHeight("0").
		ChildWidth("0").
		Items("[]").
		List("false").
		MinWidth("0").
		ReuseItem("false").
		ScrollBarFadeTime("500").
		ScrollBarMinLength("20").
		ScrollBarWidth("6").
		ScrollBarMargin("0").
		Scrollable("true").
		Virtual("false").
		HandleItemCompute("undefined")
	return ret
}

type ContainerOpt struct {
	*BaseOpt[ContainerOpt]
	handleItemUpdate  string
	handleItemClick   string
	handleItemCompute string
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
func (o *ContainerOpt) Scrollable(s string) *ContainerOpt       { o.SetProperty("scrollable", s); return o }
func (o *ContainerOpt) Virtual(s string) *ContainerOpt          { o.SetProperty("virtual", s); return o }
func (o *ContainerOpt) HandleItemClick(s string) *ContainerOpt  { o.handleItemClick = s; return o }
func (o *ContainerOpt) HandleItemUpdate(s string) *ContainerOpt { o.handleItemUpdate = s; return o }
func (o *ContainerOpt) HandleItemCompute(s string) *ContainerOpt {

	return o
}

func VListContainer(opt *ContainerOpt, child *Element) *Element {
	ret := Container(opt, child)
	opt.List("true").Virtual("true").Scrollable("true").Init(ret)
	return ret
}

func ListContainer(opt *ContainerOpt, child *Element) *Element {
	ret := Container(opt, child)
	opt.List("true").Scrollable("true").Init(ret)
	return ret
}

func Container(opt *ContainerOpt, child *Element) *Element {
	childOpt := NewOpt().Y("0").X("0").H("0")
	child.SetProperty("data", "null")
	if opt.handleItemUpdate != "" {
		childOpt.OnUpdated(opt.handleItemUpdate)
	}
	if opt.handleItemClick != "" {
		childOpt.OnClick(opt.handleItemClick)
	}
	childOpt.Init(child)
	ret := NewElement(ElementTypeContainer, ElementTagDiv,
		HScrollbar(NewOpt()),
		VScrollbar(NewOpt()),
		child,
	)
	opt.List("false").Virtual("false").Scrollable("false").ScrollLeft("0").ScrollTop("0").
		OnUpdated(".handleUpdated").
		OnCreated(".handleCreated").
		Init(ret)
	ret.SetMethod("handleUpdated", `function(k) {
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
	ret.SetMethod("handleCreated", `function() {
    if (!this.list) {
        const child = g.createElement(null, this.model.itemModel, this);
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
