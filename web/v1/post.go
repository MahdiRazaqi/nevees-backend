package v1

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/post"
	"github.com/neveesco/nevees-backend/user"
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
 * @apiParam {String} body post body
 * @apiParam {String} thumbnail post thumbnail
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
		UserID:    u.ID,
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
 * @api {put} /api/v1/post/:id Edit post
 * @apiVersion 1.0.0
 * @apiName editPost
 * @apiGroup Post
 *
 * @apiParam {String} title post title
 * @apiParam {String} body post body
 * @apiParam {String} thumbnail post thumbnail
 *
 * @apiSuccess {String} message success message
 *
 * @apiError {String} error error message
 */

func editPost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	id, _ := strconv.Atoi(c.Param("id"))

	formData := new(postForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		ID:        uint(id),
		Title:     formData.Title,
		Body:      formData.Body,
		Thumbnail: formData.Thumbnail,
	}
	if err := p.Update("user_id = ?", u.ID); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post updated successfully",
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
