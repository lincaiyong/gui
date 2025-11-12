package button

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/img"
	"github.com/lincaiyong/page/com/text"
)

func SourceRootButton() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		text.Text("parent.text").X("next.x2+4").Y("1").H("parent.h-2").Color(com.ColorGray110),
		img.Svg(com.SvgSourceRootFileLayer).X(".y-2").Y("parent.h/2-.h/2+1").W("8").H(".w").Color("parent.color"),
	)
	ret.W("child.w + 21").H("20").BorderRadius("3").
		BgColor(".selected ? page.theme.ComponentSelectedBgColor : page.theme.ComponentBgColor").
		Color(".selected ? page.theme.ComponentSelectedColor : page.theme.ComponentColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}

func SourceDirButton() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		text.Text("parent.text").X("2").Y("1").H("parent.h-2").Color(com.ColorGray110),
	)
	ret.W("child.w + 4").H("20").BorderRadius("3").
		BgColor(".selected ? page.theme.ComponentSelectedBgColor : page.theme.ComponentBgColor").
		Color(".selected ? page.theme.ComponentSelectedColor : page.theme.ComponentColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}

func SourceFileButton() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		text.Text("parent.text").X("next.x2+4").Y("1").H("parent.h-2").Color(com.ColorGray110),
		img.Svg("'svg/go.svg'").X("3").Y("2").W("16").H(".w"),
	)
	ret.W("child.w + 26").H("20").BorderRadius("3").
		BgColor(".selected ? page.theme.ComponentSelectedBgColor : page.theme.ComponentBgColor").
		Color(".selected ? page.theme.ComponentSelectedColor : page.theme.ComponentColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}
