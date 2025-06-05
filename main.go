package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	// Load templates
	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"Title":   "Welcome",
			"Message": "Welcome to GoWeb",
		})
	})

	app.Run(("0.0.0.0:8080"))

}
