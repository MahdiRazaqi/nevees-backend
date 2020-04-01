package v1

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	r := v1.Group("/")
	r.Use(middleware.JWT([]byte("secret-nevees")), userRequired)

	postGroup := r.Group("post")
	postGroup.POST("", addPost)
	postGroup.GET("/:id", getOnePost)
	postGroup.DELETE("/:id", removePost)

	tagGroup := r.Group("tag")
	tagGroup.POST("", addTag)
	tagGroup.GET("/:id", getOneTag)
}
