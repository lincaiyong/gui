package gui

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

func NoCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()
	}
}

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
