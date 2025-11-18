package tree

import (
	"fmt"
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/com/container"
	"github.com/lincaiyong/gui/com/div"
	"github.com/lincaiyong/gui/com/img"
	"github.com/lincaiyong/gui/com/text"
	"strconv"
)

func Tree() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		div.Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").
			BorderRadius("4").BgColor("this.focus ? g.theme.treeFocusSelectedBgColor : g.theme.treeSelectedBgColor").
			NameAs("selectedEle"),
		container.VListContainer(
			img.Svg("parent.data.collapsed ? 'svg/arrowRight.svg' : 'svg/arrowDown.svg'").NameAs("arrowEle").
				X("this.indent + parent.data.depth * 20 + 4").Y("parent.h/2-.h/2").W("17").H(".w").V("parent.data.leaf ? 0 : 1").Color(com.ColorGray110),
			img.Img("''").NameAs("iconEle").
				X("prev.x2+4").Y("parent.h/2-.h/2").W("16").H(".w"),
			text.Text("parent.data.text").X("prev.x2+4").Y("1").H("this.itemHeight - 2 * .y").Cursor("'default'"),
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
	itemHeight       com.Property `default:"22"`
	indent           com.Property `default:"0"`
	sort             com.Property `default:"true"`
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

func (b *Component) Items(s string) *Component {
	b.SetProp("items", s)
	return b
}

func (b *Component) Indent(s int) *Component {
	b.SetProp("indent", strconv.Itoa(s))
	return b
}

func (b *Component) Sort(v bool) *Component {
	b.SetProp("sort", fmt.Sprintf("%v", v))
	return b
}
