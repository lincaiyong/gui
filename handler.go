package gui

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"io/fs"
	"net/http"
	"path"
	"time"
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

// HandleRes r.GET("/res/*filepath", gui.HandleRes())
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

func HandlePage(c *gin.Context, title string, page *Element) {
	html, err := GenHtml(title, page)
	if err != nil {
		log.ErrorLog("fail to gen html: %v", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("fail to gen html: %v", err))
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// CacheMiddleware r.Use(gui.CacheMiddleware("2026-01-01 00:00:00"))
func CacheMiddleware(lastModifiedDateTime string) gin.HandlerFunc {
	if lastModifiedDateTime == "" {
		lastModifiedDateTime = "2025-01-01 00:00:00"
	}
	return func(c *gin.Context) {
		t, _ := time.Parse("2006-01-02 15:04:05", lastModifiedDateTime)
		lastModified := t.UTC().Format(http.TimeFormat)
		c.Header("Last-Modified", lastModified)
		c.Next()
	}
}

// CorsMiddleware r.Use(gui.CorsMiddleware)
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
