package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	. "github.com/lincaiyong/gui"
)

func main() {
	common.StartServer("gui", "v1.0.1", "",
		func(_ []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", HandleRes())
			r.GET("/hello", func(c *gin.Context) {
				root := Div(NewOpt(),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("10").OnClick("onText1Click"), "'hello world'").SetName("text1"),
					HDivider(NewOpt().Y("prev.y2")),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("prev.y2").OnClick("() => console.log(12)"), "'hello world'").SetName("text2"),
					Div(NewOpt().Y("450").H("400"),
						Div(NewOpt().W("next.x"),
							Editor(NewEditorOpt().X("20").Y("0").W("800").H("next.y - .y").BgColor(ColorBlue)).SetName("editor"),
							HBar(NewOpt().BgColor(ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w")),
							Div(NewOpt().X("20").Y("prev.y2").W("800").H("parent.h-prev.y2").BgColor(ColorYellow),
								ListContainer(NewContainerOpt().HandleItemCompute("onComputeItem").HandleItemUpdated("onUpdateItem"),
									Div(NewOpt().OnHover("onHoverItem"),
										Text(NewOpt(), "''"),
									),
								).SetName("container"),
							),
						),
						VBar(NewOpt().X("parent.w/2").BgColor(ColorBlue).Opacity("0.1")),
						Div(NewOpt().X("prev.x2").W("parent.w-prev.x2"),
							Compare(NewOpt().Y("0").H("next.y").BgColor(ColorRed)).SetName("compare"),
							HBar(NewOpt().BgColor(ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w")),
							Div(NewOpt().Y("prev.y2").W("40").H("40").BgColor(ColorGreen)),
							Button(NewButtonOpt().Svg(SvgProject).X("prev.x2").Y("prev.y2 + 100").W("40").H("40")),
						),
					),
				)
				HandlePage(c, "example", root, `function onText1Click() {
	console.log(...arguments);
}
function onComputeItem(containerEle, idx, prev) {
    return {
        key: ''+idx,
        x: 0,
        y: 20 * idx,
        w: 200,
        h: 20,
        text: 'hello world!' + idx,
    }
}
function onUpdateItem(k, v) {
	if (k === 'data') {
        this.child.innerText = v?.text || '';
    }
}
function onHoverItem(ele, hovered) {
    ele.backgroundColor = hovered ? '#888' : '#eee';
}`)
			})
			return nil
		})
}
