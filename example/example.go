package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/page"
	. "github.com/lincaiyong/page/com/all"
)

//go:embed example.js
var exampleJs string

func main() {
	common.StartServer("page", "v1.0.1", "",
		func(_ []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", page.HandleRes())
			r.GET("/hello", func(c *gin.Context) {
				comp := Root(
					Text("'hello world'").H("200").X("parent.w/2-.w/2").Y("100"),
					HDivider().Y("prev.y2"),
					Text("'hello world'").H("200").X("parent.w/2-.w/2").Y("prev.y2"),
				)
				page.MakePage(c, "debug", comp)
			})
			r.GET("/click", func(c *gin.Context) {
				comp := Root(
					Text("'hello world'").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").
						OnCreated("() => Root.test('onCreated')").
						OnUpdated("() => Root.test('onUpdated')").
						OnClick("Root.handleClick"),
				).Code(`
function test(msg) {
	console.log("test: " + msg);
}
function handleClick() {
	console.log(...arguments);
}`)
				page.MakePage(c, "debug", comp)
			})
			r.GET("/bar", func(c *gin.Context) {
				comp := Root(Div().SetSlots(
					Div().W("next.x").SetSlots(
						Editor().X("20").Y("0").W("800").H("next.y - .y").BgColor(ColorBlue),
						HBar().BgColor(ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w"),
						Div().X("20").Y("prev.y2").W("800").H("parent.h-prev.y2").BgColor(ColorYellow).SetSlots(
							Text("'hello world!'").H("200"),
						),
					),
					VBar().X("parent.w/2").BgColor(ColorBlue).Opacity("0.1"),
					Div().X("prev.x2").W("parent.w-prev.x2").SetSlots(
						Compare().Y("0").H("next.y").BgColor(ColorRed),
						HBar().BgColor(ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w"),
						Div().Y("prev.y2").W("40").H("40").BgColor(ColorGreen),
						Button().Icon("'svg/project.svg'").X("prev.x2").Y("prev.y2 + 100").W("40").H("40"),
					),
				))
				page.MakePage(c, "debug", comp)
			})
			r.GET("/vlist", func(c *gin.Context) {
				comp := Root(
					Div().BgColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").SetSlots(
						VListContainer(Div().OnHover("Root.hoverItem").SetSlots(
							Text("''").NameAs("textEle"),
						)).NameAs("containerEle").ItemCompute("Root.computeItem").ItemOnUpdated("Root.updateItem"),
					),
				).OnCreated("Root.onCreated").Code(exampleJs)
				page.MakePage(c, "debug5", comp)
			})
			r.GET("/container", func(c *gin.Context) {
				comp := Root(Container(Text("'hello world!'").H("400")).Scrollable(true).BgColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2"))
				page.MakePage(c, "debug4", comp)
			})
			r.GET("/editor", func(c *gin.Context) {
				comp := Root(Editor().NameAs("editorEle")).OnCreated("Root.test").
					Code(`
function test() {
	setTimeout(function() {
		const editor = page.root.editorEle;
		editor.setValue('package main\n\nfunc main() {\n\n}');
		editor.setLanguage('go');
	}, 1000);
}
`)
				page.MakePage(c, "editor", comp)
			})
			r.GET("/iframe", func(c *gin.Context) {
				comp := Root(
					Iframe().NameAs("iframeEle"),
				).OnCreated("Root.test").Code(`
function test() {
	setTimeout(function() {
		const iframe = page.root.iframeEle;
		const url = 'http://127.0.0.1:9123/editor';
		page.util.fetch(url).then(html => {
			iframe.setHtml(html);
		}).catch(e => {
			page.log.error(e);
		});
	}, 1000);
}`)
				page.MakePage(c, "iframe", comp)
			})
			r.GET("/img", func(c *gin.Context) {
				page.MakePage(c, "img", Root(Img("'img/bot.png'")))
			})
			r.GET("/input", func(c *gin.Context) {
				page.MakePage(c, "input", Root(
					Input().H("30").W("400").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").
						BorderTop(1).BorderBottom(1),
				))
			})
			r.GET("/tree", func(c *gin.Context) {
				page.MakePage(c, "tree", Root(
					Tree().NameAs("treeEle"),
				).OnCreated("Root.test").Code(`
function test() {
	setTimeout(function() {
		const treeEle = page.root.treeEle;
		treeEle.items = ['test/test.go', 'test/test.js', 'test/test.py', 'test/test.txt', 'go.mod'];
	}, 1000);
}
`))
			})
			r.GET("/toyeditor", func(c *gin.Context) {
				page.MakePage(c, "toyeditor", Root(ToyEditor()))
			})
			r.GET("/goland", goland)
			r.GET("/larkbase", larkbase)
			return nil
		},
	)
}
