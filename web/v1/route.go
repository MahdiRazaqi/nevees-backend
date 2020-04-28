package v1

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/neveesco/nevees-backend/config"
)

var signature = config.CFG.JWT.SigningKey

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	publicGroup := v1.Group("/public/post")
	publicGroup.GET("", listPosts)
	publicGroup.GET("/:id", onePost)

	r := v1.Group("/")
	r.Use(middleware.JWT([]byte(signature)), userRequired)

	postGroup := r.Group("post")
	postGroup.POST("", addPost)
	postGroup.PUT("/:id", editPost)
	postGroup.DELETE("/:id", removePost)
}
