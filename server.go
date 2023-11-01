package routix

import (
	"github.com/gin-gonic/gin"
)

type ControllerType func() *gin.Engine
type Middleware func() gin.HandlerFunc
type ServerConfig struct {
	Loader      []func()
	Controllers []ControllerType
	Middlewares []gin.HandlerFunc
	Port        string
	DebugLogger bool
}

// CreateServer creates a new Gin server with the given configuration.
// It takes a ServerConfig parameter that specifies the server's configuration.
// The function returns a *gin.Engine, which is the created Gin server.
func CreateServer(config ServerConfig) *gin.Engine {
	if !config.DebugLogger {
		gin.SetMode(gin.ReleaseMode)
	}
	// Create a new Gin server with default middleware
	Driver = gin.Default()

	// Auto apply global middlewares
	applyMiddlewares(Driver, config.Middlewares)

	// Connect the controllers to the server
	for _, controller := range config.Controllers {
		controller()
	}

	// Return the created Gin server
	return Driver
}

// applyMiddlewares applies a list of middlewares to a gin.Engine.
//
// Parameters:
//   - r: a pointer to a gin.Engine object.
//   - middlewares: a slice of gin.HandlerFunc objects.
//
// Return type: None.
func applyMiddlewares(r *gin.Engine, middlewares []gin.HandlerFunc) {
	for _, middleware := range middlewares {
		r.Use(middleware)
	}
}
