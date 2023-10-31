package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l1ttps/routix"
)

var (
	Controller   = routix.Controller
	Get          = routix.Get
	CreateServer = routix.CreateServer
)

func AppController() *gin.Engine {
	return Controller("/",
		Get("/hello-world",
			func(c *gin.Context) any {
				return "hello world"
			},
		),
	)
}

func main() {
	app := CreateServer(routix.ServerConfig{
		Controllers: []routix.ControllerType{
			AppController,
		},
		DebugLogger: true,
	})
	app.Run(":8080")
}
