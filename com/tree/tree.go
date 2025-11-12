package tree

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/container"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/img"
	"github.com/lincaiyong/page/com/text"
)

func Tree() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		div.Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").
			BorderRadius("4").BgColor("this.focus ? page.theme.treeFocusSelectedBgColor : page.theme.treeSelectedBgColor").
			NameAs("selectedEle"),
		container.VListContainer(
			img.Img("parent.data.collapsed ? 'svg/arrow-right.svg' : 'svg/arrow-down.svg'").NameAs("arrowEle").
				X("parent.data.depth * 20 + 4").Y("5").W("11").H(".w").V("parent.data.leaf ? 0 : 1"),
			img.Img("''").NameAs("iconEle").
				X("next.x-18").Y("5").W("14").H(".w"),
			text.Text("parent.data.text").X("parent.data.depth * 20 + 40").Y("2").H("this.itemHeight - 2 * .y").Cursor("'default'"),
		).Align("'fill'").X("10").W("parent.w - .x").
			NameAs("containerEle").
			ItemCompute("Tree.computeItem").
			ItemOnClick("Tree.clickItem").
			ItemOnUpdated("Tree.updateItem"),
	)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	focus            com.Property `default:"false"`
	items            com.Property `default:"[]"`
	nodeMap          com.Property `default:"undefined"`
	onClickItem      com.Property `default:"undefined"`
	selectedChildTop com.Property `default:"0"`
	itemHeight       com.Property `default:"24"`
	computeItem      com.Method   `static:"true"`
	clickItem        com.Method   `static:"true"`
	updateItem       com.Method   `static:"true"`
	onUpdated        com.Method
	makeNodeMap      com.Method
	nodeToItems      com.Method
	selectChild      com.Method
}

func (b *Component) OnClickItem(s string) *Component {
	b.SetProp("onClickItem", s)
	return b
}
