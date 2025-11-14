package main

import (
	_ "embed"
	"github.com/lincaiyong/gui"
	. "github.com/lincaiyong/gui/com"
	. "github.com/lincaiyong/gui/com/all"
	"github.com/lincaiyong/gui/com/root"
	"github.com/lincaiyong/log"
	"os"
)

//go:embed frontend.js
var frontendJs string

func main() {
	gui.SetBaseUrl("")
	comp := Root(
		Div().SetSlots(
			Div().NameAs("headerEle").H("1").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235).SetSlots(
				ToolButton().Svg(SvgSettings).X("parent.w-.w-.y-1").Y("parent.h/2-.h/2-0.5").Flag(),
			),
			Div().Y("prev.y2").H("parent.h-next.h-prev.h").SetSlots(
				Div().NameAs("leftSideEle").W("33").BgColor(ColorGray247).BorderRight(1).BorderColor(ColorGray235).SetSlots(
					ToolButton().Svg(SvgProject).X("parent.w/2-.w/2-0.5").Y(".x").OnClick("Root.onOpenProject"),
					ToolButton().Svg(SvgSearch).X("prev.x").Y("prev.y2 + 8").OnClick("Root.test"),
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
							Div().H("32").SetSlots(
								Text("'Project'").X(".y+4").Y("parent.h/2-.h/2").H("20").FontWeight("600"),
								Svg(SvgArrowDown).X("prev.x2+2").Y("parent.h/2-.h/2").W("17").H(".w").Color(ColorGray110),
							),
							Div().Y("prev.y2").H("24").SetSlots(
								Svg(SvgArrowDown).X(".y+6").Y("parent.h/2-.h/2").W("17").H(".w").Color(ColorGray110),
								Svg(SvgFolder).X("prev.x2+4").Y("parent.h/2-.h/2").W("16").H(".w").Color(ColorGray110),
								Text("root.projectName").X("prev.x2+4").Y("parent.h/2-.h/2").H("20").FontWeight("600"),
							),
							Tree().NameAs("treeEle").Y("prev.y2").H("parent.h-.y").OnClickItem("Root.onClickTreeItem").Items("root.files").Indent(16),
						),
						VBar().X("parent.w/3").BgColor(ColorYellow).Opacity("0"),
						Div().NameAs("mainPaneEle").X("prev.x2-prev.w/2").W("parent.w-.x").SetSlots(
							Editor().NameAs("editorEle").OnCursorPositionChange("Root.onCursorPositionChange").
								Value("root.currentFileContent").Language("root.currentFileLanguage"),
						),
					),
					HBar().Y("parent.h*3/5"),
					Div().NameAs("bottomPaneEle").Y("prev.y2-prev.h/2").H("parent.h-.y").BorderTop(1).BorderColor(ColorGray235).SetSlots(
						Div().NameAs("bottomHeaderEle").H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
						Div().Y("prev.y2").H("parent.h-.y").BgColor(ColorWhite).SetSlots(
							Editor().NameAs("outputEle").ShowLineNo(false).Value("root.outputText"),
						),
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
	).Code(frontendJs).OnCreated("Root.OnCreated").OnUpdated("Root.onUpdated")
	root.AddProp("files", "[]")
	root.AddProp("currentFile", "''")
	root.AddProp("sourceRoot", "''")
	root.AddProp("projectName", "''")
	root.AddProp("currentFileContent", "''")
	root.AddProp("currentFileLanguage", "'go'")
	root.AddProp("outputText", "'hello world'")
	html, err := gui.MakeHtml("CodeEdge", comp)
	if err != nil {
		log.ErrorLog("%v", err)
		return
	}
	_ = os.Mkdir("./dist", os.ModePerm)
	err = os.WriteFile("./dist/index.html", []byte(html), 0644)
	if err != nil {
		log.ErrorLog("%v", err)
		return
	}
	log.InfoLog("done")
}
