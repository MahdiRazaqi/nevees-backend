package v1

import (
	"strconv"

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

/**
 * @api {get} /api/v1/public/post List posts
 * @apiVersion 1.0.0
 * @apiName listPosts
 * @apiGroup Post
 *
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 *
 * @apiSuccess {[]Object} post array of posts model
 *
 * @apiError {String} error error message
 */

func listPosts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	posts, err := post.Find(limit, page, "", "")
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"posts": posts,
	})
}
