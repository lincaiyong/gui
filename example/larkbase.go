package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func larkbase(c *gin.Context) {
	c.String(http.StatusOK, "larkbase")
}
