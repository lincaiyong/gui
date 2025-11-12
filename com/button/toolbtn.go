package button

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/img"
)

func ToolButton() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		img.Svg("parent.svg").X("4").Y(".x").W("parent.w - 2 * .x").H(".w").Color("parent.color"),
		div.Div().X("prev.x2 - .w + 1").Y("prev.y - 1").W("8").H(".w").V("parent.flag").
			BorderRadius("4").BgColor(com.ColorOrange).
			BorderColor(com.ColorWhite).BorderLeft(1).BorderRight(1).BorderTop(1).BorderBottom(1),
	)
	ret.W("24").H(".w").BorderRadius("6").
		BgColor(".selected ? page.theme.ComponentSelectedBgColor : page.theme.ComponentBgColor").
		Color(".selected ? page.theme.ComponentSelectedColor : page.theme.ComponentColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}
