package main

import (
	"github.com/l1ttps/routix"
	"github.com/l1ttps/routix/example/boilerplate/controllers"
)

var (
	CreateServer = routix.CreateServer
)

type (
	ControllerType = routix.ControllerType
	ServerConfig   = routix.ServerConfig
)

func main() {
	CreateServer(ServerConfig{
		Controllers: []ControllerType{
			controllers.AppController,
			controllers.RenderController,
		},
		BaseViewDir: "views/*",
	}).Run(":3000")
}
