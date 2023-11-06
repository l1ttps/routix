package guards

import (
	"github.com/gin-gonic/gin"
)

func ProtectedGuard(c *gin.Context) bool {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" { // improve this
		return false
	}
	return true
}
