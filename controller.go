package routix

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1ttps/routix/exception"
)

type Engine *gin.Engine

// Controller initializes and configures a new gin.Engine instance for handling HTTP requests.
//
// The function takes a basePath string as the base path for all routes, and a variadic parameter of
// RouteBase structs representing the routes to be added to the engine.
//
// The function returns a *gin.Engine instance that has been configured with the specified routes.
func Controller(basePath string, routes ...RouteBase) *gin.Engine {
	var drive = Driver
	c := drive.Group(basePath)
	methodMap := map[HTTPMethod]func(string, ...gin.HandlerFunc) gin.IRoutes{
		"GET":    c.GET,
		"POST":   c.POST,
		"PUT":    c.PUT,
		"DELETE": c.DELETE,
		"PATCH":  c.PATCH,
	}

	for _, route := range routes {
		fmt.Println(route)
		handlerFunc, exists := methodMap[route.method]
		if !exists {
			fmt.Printf("Invalid HTTP method: %s\n", route.method)
			continue
		}
		// handlerFunc(route.basePath, append(route.middlewares, route.handler)...)
		handlerFunc(route.basePath, append(route.middlewares, PipeResponse(route.handler))...)
	}

	return drive
}

// PipeResponse is a function that takes a handler function as input and returns a gin.HandlerFunc.
//
// The handler function is responsible for processing a gin.Context and returning a response.
// The function checks the type of the response:
// - If the response is an HttpExceptionResponse, it returns a JSON response with the status code and message.
// - If the response is a map[string]interface{}, it extracts the status code and message from the map and returns a JSON response.
// - Otherwise, it returns a JSON response with the response itself.
func PipeResponse(handler func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := handler(ctx)

		if httpException, ok := response.(exception.HttpExceptionResponse); ok {
			ctx.JSON(httpException.Status, gin.H{
				"status":  httpException.Status,
				"message": httpException.Message,
			})
			return
		}

		if status, ok := response.(map[string]interface{}); ok {
			statusCode, exists := status["status"].(int)
			message, messageExists := status["message"].(string)
			if exists && messageExists {
				ctx.JSON(statusCode, gin.H{
					"status":  statusCode,
					"message": message,
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, response)
	}
}

type RouteBase struct {
	basePath    string
	handler     func(c *gin.Context) any
	method      HTTPMethod
	middlewares []gin.HandlerFunc
}

type MethodHandlerConfigs struct {
	openApi OpenApiConfigs
}

type OpenApiConfigs struct {
	title       string
	description string
	version     string
	tags        []string
	success     string
}

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

// NewRouteBase creates a new RouteBase struct with the given parameters.
//
// basePath: the base path for the route.
// handler: the handler function for the route.
// method: the HTTP method for the route.
// middlewares: an array of middleware functions for the route.
// Returns: a RouteBase struct.
func NewRouteBase(basePath string, handler func(c *gin.Context) any, method HTTPMethod, middlewares []gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath:    basePath,
		handler:     handler,
		method:      method,
		middlewares: middlewares,
	}
}

func TestMethod(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, GET, middlewares)
}

// Get returns a new RouteBase with the given base path, handler function, and
// optional middlewares.
//
// Parameters:
//   - basePath: a string representing the base path for the route
//   - handler: a gin.HandlerFunc representing the handler function for the route
//   - middlewares: a variadic list of gin.HandlerFunc representing the optional
//     middlewares for the route
//
// Return:
// - RouteBase: a new instance of RouteBase
func Get(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, GET, middlewares)
}

// Get returns a new RouteBase with the given base path, handler function, and
// optional middlewares.
//
// Parameters:
//   - basePath: a string representing the base path for the route
//   - handler: a gin.HandlerFunc representing the handler function for the route
//   - middlewares: a variadic list of gin.HandlerFunc representing the optional
//     middlewares for the route
//
// Return:
// - RouteBase: a new instance of RouteBase
func Post(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, POST, middlewares)
}

// Put creates a new RouteBase with the specified base path, handler function, and optional middlewares.
//
// Parameters:
// - basePath: the base path for the route.
// - handler: the handler function for the route.
// - middlewares: optional middleware functions to be applied to the route.
//
// Return:
// - RouteBase: a new RouteBase instance.
func Put(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PUT, middlewares)
}

// Delete returns a new RouteBase with the given base path, handler, DELETE method, and optional middlewares.
//
// Parameters:
// - basePath: The base path for the route.
// - handler: The handler function to be executed for the route.
// - middlewares: Optional middleware functions to be executed before the handler.
//
// Return:
// - RouteBase: The newly created RouteBase.
func Delete(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, DELETE, middlewares)
}

// Patch creates a new RouteBase with the specified base path, handler, and middlewares for the PATCH HTTP method.
//
// Parameters:
// - basePath: The base path for the route.
// - handler: The handler function for the route.
// - middlewares: Optional middlewares to be applied to the route.
//
// Return:
// - RouteBase: The created RouteBase object.
func Patch(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PATCH, middlewares)
}
