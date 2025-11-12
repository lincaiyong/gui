package button

import (
	"fmt"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/img"
)

func Button() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		img.Svg("parent.svg").X("4").Y(".x").W("parent.w - 2 * .x").H(".w").Color("parent.color"),
		div.Div().X("prev.x2 - .w + 1").Y("prev.y - 1").W("6").H(".w").V("0").BorderRadius("3"),
	)
	ret.W("24").H(".w").BorderRadius("6").
		BgColor(".selected ? page.theme.ComponentSelectedBgColor : page.theme.ComponentBgColor").
		Color(".selected ? page.theme.ComponentSelectedColor : page.theme.ComponentColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	svg          com.Property `default:"'svg/project.svg'"`
	selected     com.Property `default:"false"`
	text         com.Property `default:"''"`
	img          com.Property `default:"''"`
	flag         com.Property `default:"0"`
	handleHover  com.Method
	handleActive com.Method
}

func (b *Component) Svg(s string) *Component {
	b.SetProp("svg", s)
	return b
}

func (b *Component) Selected(v bool) *Component {
	b.SetProp("selected", fmt.Sprintf("%v", v))
	return b
}

func (b *Component) Text(s string) *Component {
	b.SetProp("text", s)
	return b
}

func (b *Component) Flag() *Component {
	b.SetProp("flag", "1")
	return b
}
