package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/gui"
	. "github.com/lincaiyong/gui"
)

func main() {
	common.StartServer("gui", "v1.0.1", "",
		func(_ []string, r *gin.RouterGroup) error {
			r.Use(CorsMiddleware())
			r.GET("/res/*filepath", gui.HandleRes())
			r.GET("/hello", func(c *gin.Context) {
				comp := Div(NewOpt(),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("100"), "'hello world'"),
					HDivider(NewOpt().Y("prev.y2")),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("prev.y2"), "'hello world'"),
				)
				gui.HandlePage(c, "hello", comp)
			})
			r.GET("/click", func(c *gin.Context) {
				root := Div(NewOpt(),
					Text(NewOpt().OnClick("handleClick").OnCreated("() => test('text created')"), "'hello world'"),
				)
				gui.HandlePage(c, "click", root, `function test(msg) {
	console.log("test: " + msg);
}
function handleClick() {
	console.log(...arguments);
}`)
			})
			r.GET("/container", func(c *gin.Context) {
				root := Div(NewOpt(),
					Container(NewContainerOpt().H("400").Scrollable("true").BgColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2"),
						Named("text", Text(NewOpt().W("1000").H("600"), "'hello world!'"))),
				)
				gui.HandlePage(c, "container", root)
			})
			r.GET("/editor", func(c *gin.Context) {
				root := Editor(NewEditorOpt().OnCreated("test"))
				gui.HandlePage(c, "editor", root, `function test() {
	setTimeout(function() {
		const editor = g.root.editorEle;
		editor.setValue('package main\n\nfunc main() {\n\n}');
		editor.setLanguage('go');
	}, 1000);
}`)
			})
			r.GET("/iframe", func(c *gin.Context) {
				comp := Named("iframe", Iframe(NewOpt().OnCreated("test")))
				gui.HandlePage(c, "iframe", comp, `function test() {
	setTimeout(function() {
		const iframe = g.root.iframeEle;
		const url = 'http://127.0.0.1:9123/editor';
		g.fetch(url).then(html => {
			iframe.srcdoc = html;
		}).catch(e => {
			console.error(e);
		});
	}, 1000);
}`)
			})
			r.GET("/img", func(c *gin.Context) {
				gui.HandlePage(c, "img", Img(NewOpt(), "'res/img/bot.png'"))
			})
			r.GET("/input", func(c *gin.Context) {
				gui.HandlePage(c, "input", Div(NewOpt(), Input(NewOpt().H("30").W("400").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").
					BorderTop(1).BorderBottom(1), "'x'")))
			})
			r.GET("/tree", func(c *gin.Context) {
				gui.HandlePage(c, "tree", Div(NewOpt().OnCreated("test"),
					Named("tree", Tree(NewTreeOpt().OnClickItem("handleClickItem"))),
				), `
function handleClickItem(ele) {
	console.log(ele);
}
function test() {
	setTimeout(function() {
		const treeEle = g.root.treeEle;
		treeEle.items = ['test/test.go', 'test/test.js', 'test/test.py', 'test/test.txt', 'go.mod'];
	}, 1000);
}
`)
			})
			return nil
		},
	)
}
