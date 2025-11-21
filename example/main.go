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
				comp := Div(NewOpt(),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("100"), "'hello world'"),
					HDivider(NewOpt().Y("prev.y2")),
					Text(NewOpt().H("200").X("parent.w/2-.w/2").Y("prev.y2"), "'hello world'"),
				)
				HandlePage(c, "example", comp)
			})
			return nil
		})
}
