package controllers

import (
	"github.com/gin-gonic/gin"
)

func RenderController() {
	Controller("render",
		Get("/hello-world", func(c *gin.Context) any {
			return gin.H{
				"message": "Hello World",
			}
		}, Render("index.tmpl")),
	)
}
