package main

import (
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/gui/com"
	. "github.com/lincaiyong/gui/com/all"
	"github.com/lincaiyong/log"
	"os"
)

func main() {
	com.BaseUrl = ""
	comp := Root(
		Div().H("parent.h-next.h").SetSlots(
			Div().NameAs("leftSideEle").W("32").SetSlots(
				Button().OnClick("Root.handleClick"),
			),
			Div().X("prev.x2").W("parent.w-prev.w-next.w").BgColor("g.theme.grayPaneColor").SetSlots(
				Div().NameAs("navEle").W("next.x").SetSlots(
					Tree().NameAs("treeEle").OnClickItem("Root.clickTreeItem"),
				),
				VBar().X("parent.w/3").BgColor(ColorYellow).Opacity("0.1"),
				Div().NameAs("mainEle").X("prev.x2").W("parent.w-prev.x2").SetSlots(
					Editor().NameAs("editorEle"),
				),
			),
			Div().NameAs("rightSideEle").X("parent.w-.w").W("32").BgColor("g.theme.grayPaneColor"),
		),
		Div().NameAs("footerEle").Y("parent.h-.h").H("24").BgColor("g.theme.grayPaneColor"),
	).OnCreated("Root.test").
		Code(`
function clickTreeItem(itemEle) {
	Root.log('click: ' + JSON.stringify(itemEle.data));
	if (itemEle.data.leaf) {
		const relPath = itemEle.data.key;
		Root.readFile(g.state.folder + '/' + relPath).then(v => {
			g.root.editorEle.setValue(v);
		});
	}
}
function handleClick() {
	Root.selectFolder().then(s => {
		const obj = JSON.parse(s)
		Root.log(obj);
		g.state.folder = obj.folder; 
		g.root.treeEle.items = obj.files;
	});
}
function log(v) {
	go.main.App.Log(v);
}
function selectFolder() {
	return go.main.App.SelectFolder();
}
function readFile(path) {
	return go.main.App.ReadFile(path);
}
`)
	html, err := page.MakeHtml("CodeEdge", comp)
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
