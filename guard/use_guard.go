package guard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UseGuard is a function that takes in one or more authentication functions and returns a Gin middleware handler.
//
// The authentication functions are passed in as variadic arguments, represented by the `authFuncs` parameter. These functions take in a Gin context (`c *gin.Context`) and return a boolean value indicating whether the authentication is successful or not.
//
// The middleware handler returned by UseGuard iterates over each authentication function in the `authFuncs` slice. If any of the authentication functions return `false`, indicating that the authentication is unsuccessful, the handler responds with a JSON error message and aborts the request.
//
// If all authentication functions return `true`, indicating that the authentication is successful, the handler calls the `Next()` method on the Gin context to pass the request to the next middleware or route handler in the chain.
//
// UseGuard does not have any return values.
func UseGuard(authFuncs ...func(c *gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, authFunc := range authFuncs {
			if !authFunc(c) {
				statusCode := http.StatusForbidden
				c.JSON(statusCode, gin.H{
					"status":  statusCode,
					"message": "Forbidden",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
