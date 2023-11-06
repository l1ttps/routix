package interceptor

import (
	"github.com/gin-gonic/gin"
)

type InterceptorContext struct {
	*gin.Context
}

// CallHandler is a function type that represents the next function to be called.
type CallHandler func()

// UseInterceptor is a function that takes an interceptor function and returns a gin.HandlerFunc.
//
// The interceptor function takes an InterceptorContext pointer as a parameter and returns a function.
// The InterceptorContext encapsulates the *gin.Context object.
// The returned function is the next handler to be executed in the middleware chain.
//
// The UseInterceptor function itself returns a gin.HandlerFunc, which is a middleware function that conforms to the gin middleware signature.
// It takes a *gin.Context object as a parameter, creates a new InterceptorContext, and invokes the interceptor function with the InterceptorContext.
// It then calls the Next() method on the *gin.Context to proceed to the next middleware.
// Finally, it executes the nextHandler function returned by the interceptor function to execute the interceptor logic after the request is handled by other middlewares.
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
