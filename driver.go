package routix

import "github.com/gin-gonic/gin"

var Driver *gin.Engine

var PathRoot string = "/"

var BaseViewDir string = "views/*"

var IsEnableRender bool = false
