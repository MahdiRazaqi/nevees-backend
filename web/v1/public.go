package v1

import (
	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/post"
)

/**
 * @api {get} /api/v1/public/post/:id Get a post
 * @apiVersion 1.0.0
 * @apiName onePost
 * @apiGroup Post
 *
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error error message
 */

func onePost(c echo.Context) error {
	id := c.Param("id")

	p := &post.Post{}
	if err := p.FindOne("id = ?", id); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"post": p,
	})
}
