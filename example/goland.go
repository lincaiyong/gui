package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/page"
	. "github.com/lincaiyong/page/com/all"
)

func goland(c *gin.Context) {
	comp := Root(
		Div().W("2782/2").H("1590/2").SetSlots(
			Div().NameAs("headerEle").H("32").BgColor(ColorYellow).Opacity("0.1"),
			Div().Y("prev.y2").H("parent.h-next.h-prev.h").SetSlots(
				Div().NameAs("leftSideEle").W("32").BgColor(ColorCyan).SetSlots(
					Button().OnClick("Root.handleClick").Icon("'svg/project.svg'"),
				),
				Div().X("prev.x2").W("parent.w-prev.w-next.w").BgColor("page.theme.grayPaneColor").SetSlots(
					Div().NameAs("navEle").W("next.x").SetSlots(
						Tree().NameAs("treeEle").OnClickItem("Root.clickTreeItem"),
					),
					VBar().X("parent.w/3").BgColor(ColorYellow).Opacity("0.1"),
					Div().NameAs("mainEle").X("prev.x2").W("parent.w-prev.x2").SetSlots(
						Editor().NameAs("editorEle"),
					),
				),
				Div().NameAs("rightSideEle").X("parent.w-.w").W("32").BgColor(ColorCyan).Opacity("0.1"),
			),
			Div().NameAs("footerEle").Y("parent.h-.h").H("24").BgColor(ColorYellow).Opacity("0.1"),
			Img("'img/goland.png'").NameAs("imgEle"),
		),
		Button().OnClick("Root.handleClick").Y("prev.y2").X("parent.w/2-.w/2"),
		Button().OnClick("Root.handleClick2").Y("prev.y2").X("parent.w/2-.w/2"),
	).Code(`
function handleClick() {
	const img = page.root.imgEle;
	img.v = !img.v;
}

function handleClick2() {
	const img = page.root.imgEle;
	img.opacity = img.opacity >= 1 ? 0.4 : img.opacity + 0.3;
}`)
	page.MakePage(c, "goland", comp)
}
