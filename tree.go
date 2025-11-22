package gui

import "fmt"

func NewTreeOpt() *TreeOpt {
	ret := &TreeOpt{}
	ret.BaseOpt = NewBaseOpt[TreeOpt](ret)
	ret.Focus("false").
		Items("[]").
		NodeMap("undefined").
		OnClickItem("undefined").
		SelectedChildTop("0").
		ItemHeight("22").
		Indent("0").
		Sort("true")
	return ret
}

type TreeOpt struct {
	*BaseOpt[TreeOpt]
}

func (o *TreeOpt) Focus(s string) *TreeOpt            { o.SetProperty("focus", s); return o }
func (o *TreeOpt) Items(s string) *TreeOpt            { o.SetProperty("items", s); return o }
func (o *TreeOpt) NodeMap(s string) *TreeOpt          { o.SetProperty("nodeMap", s); return o }
func (o *TreeOpt) OnClickItem(s string) *TreeOpt      { o.SetProperty("onClickItem", s); return o }
func (o *TreeOpt) SelectedChildTop(s string) *TreeOpt { o.SetProperty("selectedChildTop", s); return o }
func (o *TreeOpt) ItemHeight(s string) *TreeOpt       { o.SetProperty("itemHeight", s); return o }
func (o *TreeOpt) Indent(s string) *TreeOpt           { o.SetProperty("indent", s); return o }
func (o *TreeOpt) Sort(s string) *TreeOpt             { o.SetProperty("sort", s); return o }

const treeFocusSelectedBgColor = "'#D5E1FF'"
const treeSelectedBgColor = "'#DFE1E5'"

func Tree(opt *TreeOpt) *Element {
	treeItem := Div(NewOpt(),
		Named("arrow", Svg(NewOpt().X("parent.data.depth * 20 + 4").Y("parent.h/2-.h/2").W("17").H(".w").
			V("parent.data.leaf ? 0 : 1").Color(ColorGray110),
			fmt.Sprintf("parent.data.collapsed ? %s : %s", SvgArrowRight, SvgArrowDown))),
		Named("icon", Img(NewOpt().X("prev.x2+4").Y("parent.h/2-.h/2").W("16").H(".w"), "''")),
		Text(NewOpt().X("prev.x2+4").Y("1").H("parent.h - 2 * .y").Cursor("'default'"), "parent.data.text"),
	).SetLocalRoot()
	ret := NewElement(ElementTypeTree, ElementTagDiv,
		Named("selected", Div(NewOpt().X("10").Y("local.selectedChildTop-next.scrollTop").W("parent.w-20").H("local.itemHeight").
			BorderRadius("4").BgColor(fmt.Sprintf("local.focus ? %s : %s", treeFocusSelectedBgColor, treeSelectedBgColor)))),
		Named("container", VListContainer(
			NewContainerOpt().Align("'fill'").X("10").W("parent.w - .x").
				HandleItemCompute("treeItem_handleCompute").
				HandleItemClick("treeItem_handleClick").
				HandleItemUpdate("treeItem_handleUpdated"),
			treeItem,
		)),
	)
	ret.SetLocalRoot()
	opt.OnUpdated("tree_handleUpdated").Init(ret)
	return ret
}
