package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	. "github.com/lincaiyong/gui/com"
	. "github.com/lincaiyong/gui/com/all"
	"github.com/lincaiyong/gui/com/root"
)

func goland(c *gin.Context) {
	comp := Root(
		Div().W("2782/2").H("1590/2").SetSlots(
			Div().NameAs("headerEle").H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235).SetSlots(
				ToolButton().Svg(SvgSettings).X("parent.w-.w-.y-1").Y("parent.h/2-.h/2-0.5").Flag(),
			),
			Div().Y("prev.y2").H("parent.h-next.h-prev.h").SetSlots(
				Div().NameAs("leftSideEle").W("33").BgColor(ColorGray247).BorderRight(1).BorderColor(ColorGray235).SetSlots(
					ToolButton().Svg(SvgProject).X("parent.w/2-.w/2-0.5").Y(".x").OnClick("Root.handleClick"),
					ToolButton().Svg(SvgCommit).X("prev.x").Y("prev.y2 + 8"),
					ToolButton().Svg(SvgPullRequests).X("prev.x").Y("prev.y2 + 8"),
					HDivider().X("prev.x").Y("prev.y2 + 9").W("prev.w").BgColor(ColorGray201),
					ToolButton().Svg(SvgStructure).X("prev.x").Y("prev.y2 + 9"),
					ToolButton().Svg(SvgMoreHorizontal).X("prev.x").Y("prev.y2 + 8"),

					ToolButton().Svg(SvgPythonPackages).X("next.x").Y("next.y-8-.h"),
					ToolButton().Svg(SvgServices).X("next.x").Y("next.y-8-.h"),
					ToolButton().Svg(SvgTerminal).X("next.x").Y("next.y-8-.h"),
					ToolButton().Svg(SvgProblems).X("next.x").Y("next.y-8-.h"),
					ToolButton().Svg(SvgVCS).X("parent.w/2-.w/2-0.5").Y("parent.h-.h-.x"),
				),
				Div().X("prev.x2").W("parent.w-prev.w-next.w").BgColor(ColorGray247).SetSlots(
					Div().H("next.y+next.h/2").SetSlots(
						Div().NameAs("leftPaneEle").W("next.x+next.w/2").BorderColor(ColorGray235).BorderRight(1).SetSlots(
							Div().H("33").SetSlots(
								Text("'Project'").X(".y+4").Y("parent.h/2-.h/2").H("18").FontWeight("600"),
								Svg(SvgArrowDown).X("prev.x2+2").Y("parent.h/2-.h/2").W("17").H(".w").Color(ColorGray110),
							),
							Tree().NameAs("treeEle").Y("prev.y2").H("parent.h-.y").OnClickItem("Root.clickTreeItem").Items("root.treeItems"),
						),
						VBar().X("parent.w/3").BgColor(ColorYellow).Opacity("0"),
						Div().NameAs("mainPaneEle").X("prev.x2-prev.w/2").W("parent.w-.x").SetSlots(
							Editor().NameAs("editorEle"),
						),
					),
					HBar().Y("parent.h/2"),
					Div().NameAs("bottomPaneEle").Y("prev.y2-prev.h/2").H("parent.h-.y").BorderTop(1).BorderColor(ColorGray235).SetSlots(
						Div().NameAs("bottomHeaderEle").H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
						Div().Y("prev.y2").H("parent.h-.y").BgColor(ColorWhite),
					),
				),
				Div().NameAs("rightSideEle").X("parent.w-.w").W("33").BgColor(ColorGray247).BorderColor(ColorGray235).BorderLeft(1).SetSlots(
					ToolButton().Svg(SvgNotifications).X("parent.w/2-.w/2-0.5").Y(".x").Flag(),
					ToolButton().Svg(SvgAIChat).X("prev.x").Y("prev.y2 + 8"),
					ToolButton().Svg(SvgDatabase).X("prev.x").Y("prev.y2 + 8"),
					ToolButton().Svg(SvgPythonConsole).X("prev.x").Y("prev.y2 + 8"),
				),
			),
			Div().NameAs("footerEle").Y("parent.h-.h").H("24").BgColor(ColorGray247).BorderColor(ColorGray235).BorderTop(1).SetSlots(
				SourceRootButton().Text("'page'").X("9").Y("1"),
				Svg(SvgArrowRight).X("prev.x2").Y("parent.h/2-.h/2-1").W("17").H(".w").Color(ColorGray110),
				SourceDirButton().Text("'example'").X("prev.x2 + 1").Y("1"),
				Svg(SvgArrowRight).X("prev.x2").Y("parent.h/2-.h/2-1").W("17").H(".w").Color(ColorGray110),
				SourceFileButton().Text("'goland.go'").X("prev.x2 + 1").Y("1"),
			),
			Img("'img/goland.png'").NameAs("imgEle").V("0"),
		),
		Button().OnClick("Root.handleClick").Y("prev.y2").X("parent.w/2-.w/2"),
		Button().OnClick("Root.handleClick2").Y("prev.y2").X("parent.w/2-.w/2"),
	).Code(`
function handleClick() {
	const img = g.root.imgEle;
	img.v = !img.v;
}

function handleClick2() {
	const img = g.root.imgEle;
	img.opacity = img.opacity >= 1 ? 0.4 : img.opacity + 0.3;
}

function clickTreeItem(itemEle) {
	console.log(itemEle.data);
	if (itemEle.data.leaf) {
		g.root.currentFile = itemEle.data;
		g.root.editorEle.setValue(itemEle.data.key);
	}
}

function startup() {
	setTimeout(function() {
		g.root.treeItems = [
			'page/com/test.go',
			'page/example/goland/explorer_pane.js',
			'page/example/goland/explorer_pane.jsm',
			'page/example/goland/icons.txt',
			'page/example/goland/top.js',
			'page/example/goland/top.jsm',
			'page/example/example.go',
			'page/example/example.js',
			'page/example/go.mod',
			'page/example/goland.go',
			'page/example/larkbase.go',
			'page/example/main.go',
			'page/example/page.log',
		];
	});
}`).OnCreated("Root.startup")
	root.AddProp("treeItems", "[]")
	root.AddProp("currentFile", "''")
	root.AddProp("sourceRoot", "'/Users/bytedance/Code/lincaiyong'")
	gui.MakePage(c, "goland", comp)
}
