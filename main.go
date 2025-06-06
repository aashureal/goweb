package main

import (
	"net/http"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Welcome",
			"message": "Welcome to Goweb",
		})
	})

	// Port for Render
	port := "10000"
	if p := getenv("PORT"); p != "" {
		port = p
	}
	r.Run(":" + port)
}

// getenv returns the environment variable or empty string if not found
func getenv(key string) string {
	val, ok := syscall.Getenv(key)
	if ok {
		return val
	}
	return ""
}
