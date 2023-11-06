package routix

import (
	"fmt"
	"net/http"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/l1ttps/routix/logger"
)

type ServerConfig struct {
	Controllers []ControllerType
	Middlewares []gin.HandlerFunc
	DebugLogger bool
	PathRoot    string
	BaseViewDir string
}

// CreateServer creates a new Gin server with the given configuration.
// It takes a ServerConfig parameter that specifies the server's configuration.
// The function returns a *gin.Engine, which is the created Gin server.
func CreateServer(config ServerConfig) *gin.Engine {

	// Check if debug logger is enabled
	if !config.DebugLogger {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new Gin server with default middleware
	Driver = gin.Default()

	// Set default base view dir ("/views")

	// Auto apply global middlewares
	applyMiddlewares(Driver, config.Middlewares)

	// Apply base path
	if config.PathRoot != "" && config.PathRoot != "/" {
		PathRoot = config.PathRoot
	}

	// Connect the controllers to the server
	connectControllers(config.Controllers)

	// Mapping views folder
	useBaseViewDir(config.BaseViewDir)

	// Fallback
	fallback()

	// Return the created Gin server
	return Driver
}

// applyMiddlewares applies a list of middlewares to a gin.Engine.
//
// Parameters:
// - r: a pointer to a gin.Engine to which the middlewares will be applied.
// - middlewares: a slice of gin.HandlerFunc representing the middlewares to be applied.
//
// Returns: nothing.
func applyMiddlewares(r *gin.Engine, middlewares []gin.HandlerFunc) {
	log := logger.Logger("Routix")

	for _, middleware := range middlewares {
		funcName := getFunctionName(middleware)
		parts := strings.Split(funcName, ".")
		log.Success(fmt.Sprintf("{%s} Applied middleware: {%s()}", parts[0], parts[1]))
		r.Use(middleware)
	}
	return
}

// connectControllers connects the given controllers and executes each controller.
//
// controllers: a slice of ControllerType that represents the controllers to be connected.
// The function iterates over the slice and executes each controller
func connectControllers(controllers []ControllerType) {
	for _, controller := range controllers {
		controller()
	}
	return
}

// UseBaseViewDir sets the base view directory for loading HTML templates.
//
// customBaseViewDir: A string representing a custom base view directory. If
// provided, the base view directory will be set to this value. Otherwise, the
// default value "views/*" will be used.
//
// No return value.
func useBaseViewDir(customBaseViewDir string) {
	if customBaseViewDir != "" && customBaseViewDir != "/" {
		BaseViewDir = customBaseViewDir
		Driver.LoadHTMLGlob(BaseViewDir)
		IsEnableRender = true
	}
	return
}

func getFunctionName(fcn interface{}) string {
	pc := reflect.ValueOf(fcn).Pointer()
	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		return ""
	}

	funcName := path.Base(funcInfo.Name())
	return funcName
}

func fallback() {
	// Fallback method not allowed
	Driver.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  http.StatusMethodNotAllowed,
			"message": http.StatusText(http.StatusMethodNotAllowed),
		})
	})

	// Fallback router not found
	Driver.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
		})
	})
}
