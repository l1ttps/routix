package routix

import "github.com/gin-gonic/gin"

const (
	RENDER string = "ROUTIX_RENDER"
)

type RenderParams struct {
	path string
}

// Render generates a gin.HandlerFunc that sets the RENDER context key to the provided path.
//
// path: The path to be set in the RENDER context key.
// Return: A gin.HandlerFunc that sets the RENDER context key.
func Render(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(RENDER, path)
	}
}
