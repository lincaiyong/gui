package gui

import (
	"embed"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"io/fs"
	"net/http"
	"path"
)

//go:embed res/**/*
var resFs embed.FS

var resFileMap map[string][]byte

func init() {
	resFileMap = make(map[string][]byte)
	err := fs.WalkDir(resFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := resFs.ReadFile(path)
		if err != nil {
			return err
		}
		path = path[4:]
		resFileMap[path] = b
		return nil
	})
	if err != nil {
		log.FatalLog("fail to walk: %v", err)
	}
}

func HandleRes() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := c.Param("filepath")[1:]
		b, ok := resFileMap[filePath]
		if !ok {
			c.String(http.StatusNotFound, "resource not found")
		}
		ext := path.Ext(filePath)
		contentType := "text/plain"
		if ext == ".css" {
			contentType = "text/css"
		} else if ext == ".js" {
			contentType = "application/javascript"
		} else if ext == ".svg" {
			contentType = "image/svg+xml"
		} else if ext == ".png" {
			contentType = "image/png"
		} else if ext == ".jpg" {
			contentType = "image/jpeg"
		}
		c.Data(http.StatusOK, contentType, b)
	}
}
