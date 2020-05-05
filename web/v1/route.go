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

	publicGroup := v1.Group("/public/")
	publicGroup.GET("post", listPosts)
	publicGroup.GET("post/:id", onePost)
	publicGroup.GET("comment", listComments)

	r := v1.Group("/")
	r.Use(middleware.JWT([]byte(signature)), userRequired)

	postGroup := r.Group("post")
	postGroup.POST("", addPost)
	postGroup.GET("", listMyPosts)
	postGroup.PUT("/:id", editPost)
	postGroup.DELETE("/:id", removePost)

	bookmarkGroup := r.Group("bookmark")
	bookmarkGroup.POST("", addToBookmark)
	bookmarkGroup.DELETE("/:id", removeFromBookmark)

	commentGroup := r.Group("comment")
	commentGroup.POST("", addComment)
}
