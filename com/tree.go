package com

func NewTreeOpt() TreeOpt {
	ret := TreeOpt{
		BaseOpt: NewBaseOpt(),
	}
	ret.SetProperty("focus", "false")
	ret.SetProperty("items", "[]")
	ret.SetProperty("nodeMap", "undefined")
	ret.SetProperty("onClickItem", "undefined")
	ret.SetProperty("selectedChildTop", "0")
	ret.SetProperty("itemHeight", "22")
	ret.SetProperty("indent", "0")
	ret.SetProperty("sort", "true")
	return ret
}

type TreeOpt struct {
	*BaseOpt
}

func (o *TreeOpt) Focus(s string) *TreeOpt            { o.SetProperty("focus", s); return o }
func (o *TreeOpt) Items(s string) *TreeOpt            { o.SetProperty("items", s); return o }
func (o *TreeOpt) NodeMap(s string) *TreeOpt          { o.SetProperty("nodeMap", s); return o }
func (o *TreeOpt) OnClickItem(s string) *TreeOpt      { o.SetProperty("onClickItem", s); return o }
func (o *TreeOpt) SelectedChildTop(s string) *TreeOpt { o.SetProperty("selectedChildTop", s); return o }
func (o *TreeOpt) ItemHeight(s string) *TreeOpt       { o.SetProperty("itemHeight", s); return o }
func (o *TreeOpt) Indent(s string) *TreeOpt           { o.SetProperty("indent", s); return o }
func (o *TreeOpt) Sort(s string) *TreeOpt             { o.SetProperty("sort", s); return o }

func Tree() *Element {
	ret := NewElement("tree", "div",
		Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").
			BorderRadius("4").BgColor("this.focus ? g.theme.treeFocusSelectedBgColor : g.theme.treeSelectedBgColor"),
		//NameAs("selectedEle"),
		VListContainer(
			//.NameAs("arrowEle").
			Svg("parent.data.collapsed ? 'svg/arrowRight.svg' : 'svg/arrowDown.svg'").
				X("this.indent + parent.data.depth * 20 + 4").Y("parent.h/2-.h/2").W("17").H(".w").V("parent.data.leaf ? 0 : 1").Color(ColorGray110),
			// NameAs("iconEle")
			Img("''").X("prev.x2+4").Y("parent.h/2-.h/2").W("16").H(".w"),
			Text("parent.data.text").X("prev.x2+4").Y("1").H("this.itemHeight - 2 * .y").Cursor("'default'"),
		).Align("'fill'").X("10").W("parent.w - .x").
			ItemCompute("Tree.computeItem").
			ItemOnClick("Tree.clickItem").
			ItemOnUpdated("Tree.updateItem"),
	)
	ret.SetLocalRoot(true)
	// NameAs("containerEle").
	ret.SetMethod("onUpdated", `function(k, v) {
    if (k === 'items') {
        this.nodeMap = tree_makeNodeMap(v, this.sort);
        this.containerEle.items = tree_nodeToItems(this.nodeMap, '', 0, 0);
        this.selectedEle.v = 0;
    }
}`)
	ret.SetMethod("selectChild", `function(child, focus) {
    this.selectedChildTop = child.y + this.containerEle.scrollTop;
    this.selectedEle.v = 1;
    this.focus = focus;
}`)
	return ret
}
