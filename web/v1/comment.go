package v1

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/comment"
	"github.com/neveesco/nevees-backend/user"
)

type commentForm struct {
	Body   string `json:"body" form:"body"`
	PostID uint   `json:"post_id" form:"post_id"`
}

/**
 * @api {post} /api/v1/comment Add comment
 * @apiVersion 1.0.0
 * @apiName addComment
 * @apiGroup Comment
 *
 * @apiParam {String} body comment body
 * @apiParam {String} post_id post id
 *
 * @apiSuccess {String} message success message
 * @apiSuccess {Object} comment comment model
 *
 * @apiError {String} error error message
 */

func addComment(c echo.Context) error {
	u := c.Get("user").(*user.User)
	formData := new(commentForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	comment := &comment.Comment{
		Body:   formData.Body,
		PostID: formData.PostID,
		UserID: u.ID,
	}
	if err := comment.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "comment created successfully",
		"comment": comment,
	})
}

/**
 * @api {get} /api/v1/public/comment List comments
 * @apiVersion 1.0.0
 * @apiName listComments
 * @apiGroup Comment
 *
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 *
 * @apiSuccess {[]Object} comments array of comment model
 *
 * @apiError {String} error error message
 */

func listComments(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	comments, err := comment.Find(limit, page, "", "")
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"comments": comments,
	})
}
