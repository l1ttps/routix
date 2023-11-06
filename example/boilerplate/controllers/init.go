package controllers

import (
	"github.com/l1ttps/routix"
	"github.com/l1ttps/routix/example/boilerplate/interceptors"
	"github.com/l1ttps/routix/guard"
	"github.com/l1ttps/routix/interceptor"
)

var (
	Get               = routix.Get
	Controller        = routix.Controller
	CreateServer      = routix.CreateServer
	UseGuard          = guard.UseGuard
	LoggerInterceptor = interceptors.LoggerInterceptor
	UseInterceptor    = interceptor.UseInterceptor
	Render            = routix.Render
)

type (
	ControllerType = routix.ControllerType
	ServerConfig   = routix.ServerConfig
)
