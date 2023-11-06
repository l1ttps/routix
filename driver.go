package routix

import "github.com/gin-gonic/gin"

// Driver is the global instance of the Gin engine
var Driver *gin.Engine

// PathRoot is the root path for the API
var PathRoot string = "/"

// BaseViewDir is the directory for the views
var BaseViewDir string = "views/*"

// IsEnableRender determines whether rendering is enabled or not
var IsEnableRender bool = false

type ControllerType func()

type MiddlewareType gin.HandlerFunc
