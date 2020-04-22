package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/post"
	"github.com/MahdiRazaqi/nevees-backend/user"
	"github.com/labstack/echo"
)

type postForm struct {
	Title     string `json:"title" form:"title"`
	Body      string `json:"body" form:"body"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}

/**
 * @api {post} /api/v1/post Add post
 * @apiVersion 1.0.0
 * @apiName addPost
 * @apiGroup Post
 *
 * @apiParam {String} title post title
 * @apiParam {String} content post content
 * @apiParam {[]String} tags post tags
 *
 * @apiSuccess {String} message success message
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error error message
 */

func addPost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	formData := new(postForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		Title:     formData.Title,
		Body:      formData.Body,
		Thumbnail: formData.Thumbnail,
		UserID:    int(u.ID),
	}
	if err := p.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post created successfully",
		"post":    p,
	})
}

/**
 * @api {delete} /api/v1/post/:id Remove post
 * @apiVersion 1.0.0
 * @apiName removePost
 * @apiGroup Post
 *
 * @apiSuccess {String} message success message
 *
 * @apiError {String} error error message
 */

func removePost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	id := c.Param("id")

	if err := post.Delete("id = ? and user_id = ?", id, u.ID); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post removed successfully",
	})
}
