package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/l1ttps/routix/example/boilerplate/guards"
)

func AppController() {
	Controller("/",
		Get("/hello-world", func(c *gin.Context) any {
			return "Hello World"
		}),

		Get("/protected",
			func(c *gin.Context) any {
				return "Hello World"
			},
			UseGuard(guards.ProtectedGuard)),

		Get("/logger",
			func(c *gin.Context) any {
				return "Logger is working"
			},
			UseInterceptor(LoggerInterceptor)),
	)
}
