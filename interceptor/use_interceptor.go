package interceptor

import (
	"github.com/gin-gonic/gin"
)

type InterceptorContext struct {
	*gin.Context
}

// CallHandler is a function type that represents the next function to be called.
type CallHandler func()

// UseInterceptor is a function that takes an interceptor function and returns a middleware function.
// The interceptor function takes an InterceptorContext and returns another function.
// The middleware function takes a *gin.Context and invokes the interceptor function, passing in an InterceptorContext.
// It then calls the Next() method on the InterceptorContext and invokes the returned function from the interceptor.

func UseInterceptor(interceptorFunc func(c *InterceptorContext) func()) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new InterceptorContext with the *gin.Context
		context := InterceptorContext{c}

		// Invoke the interceptor function and get the next handler
		nextHandler := interceptorFunc(&context)

		// Call the Next() method on the InterceptorContext to proceed to the next middleware
		c.Next()

		// Invoke the nextHandler to execute the interceptor logic after the request is handled by other middlewares
		nextHandler()
	}
}
