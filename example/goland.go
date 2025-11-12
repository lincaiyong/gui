package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/page"
	. "github.com/lincaiyong/page/com/all"
)

func goland(c *gin.Context) {
	comp := Root(
		Div().W("2782/2").H("1590/2").SetSlots(
			Div().NameAs("headerEle").H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
			Div().Y("prev.y2").H("parent.h-next.h-prev.h").SetSlots(
				Div().NameAs("leftSideEle").W("33").BgColor(ColorGray247).BorderRight(1).BorderColor(ColorGray235).SetSlots(
					Button().Icon(SvgProject).X("parent.w/2-.w/2-0.5").Y(".x").OnClick("Root.handleClick"),
					Button().Icon(SvgCommit).X("prev.x").Y("prev.y2 + 8"),
					Button().Icon(SvgPullRequests).X("prev.x").Y("prev.y2 + 8"),
					HDivider().X("prev.x").Y("prev.y2 + 9").W("prev.w").BgColor(ColorGray201),
					Button().Icon(SvgStructure).X("prev.x").Y("prev.y2 + 9"),
					Button().Icon(SvgMoreHorizontal).X("prev.x").Y("prev.y2 + 8"),

					Button().Icon(SvgPythonPackages).X("next.x").Y("next.y-8-.h"),
					Button().Icon(SvgServices).X("next.x").Y("next.y-8-.h"),
					Button().Icon(SvgTerminal).X("next.x").Y("next.y-8-.h"),
					Button().Icon(SvgProblems).X("next.x").Y("next.y-8-.h"),
					Button().Icon(SvgVCS).X("parent.w/2-.w/2-0.5").Y("parent.h-.h-.x"),
				),
				Div().X("prev.x2").W("parent.w-prev.w-next.w").BgColor(ColorGray247).SetSlots(
					Div().NameAs("navEle").W("next.x").SetSlots(
						Tree().NameAs("treeEle").OnClickItem("Root.clickTreeItem"),
					),
					VBar().X("parent.w/3").BgColor(ColorYellow),
					Div().NameAs("mainEle").X("prev.x2").W("parent.w-prev.x2").SetSlots(
						Editor().NameAs("editorEle"),
					),
				),
				Div().NameAs("rightSideEle").X("parent.w-.w").W("33").BgColor(ColorGray247).BorderColor(ColorGray235).BorderLeft(1),
			),
			Div().NameAs("footerEle").Y("parent.h-.h").H("24").BgColor(ColorGray247).BorderColor(ColorGray235).BorderTop(1).SetSlots(),
			Img("'img/goland.png'").NameAs("imgEle").V("0"),
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
