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
	o.SetProperty("computeItemFn", s)
	return o
}

func VListContainer(opt *ContainerOpt, child *Element) *Element {
	ret := Container(opt, child)
	NewOpt().W("0").H("0").Init(child)
	opt.List("true").Virtual("true").Scrollable("true").Init(ret)
	return ret
}

func ListContainer(opt *ContainerOpt, child *Element) *Element {
	ret := Container(opt, child)
	NewOpt().W("0").H("0").Init(child)
	opt.List("true").Scrollable("true").Init(ret)
	return ret
}

func Container(opt *ContainerOpt, child *Element) *Element {
	childOpt := NewOpt()
	child.SetProperty("data", "{}")
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
	opt.List("false").Virtual("false").Scrollable("true").ScrollLeft("0").ScrollTop("0").
		OnUpdated("container_handleUpdated").
		OnCreated("container_handleCreated").
		Init(ret)
	return ret
}
