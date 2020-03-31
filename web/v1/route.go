package v1

import (
	"github.com/labstack/echo"
)

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	r := v1.Group("/")
	// r.Use()

	postGroup := r.Group("post")
	postGroup.POST("", addPost)

	// e.GET("/api/v1", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "v1")
	// })

	// e.GET("/api/v1/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "v1/")
	// })
}
