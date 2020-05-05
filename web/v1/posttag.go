package v1

import (
	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/posttag"
)

type posttagForm struct {
	TagID  uint `json:"tag_id" form:"tag_id"`
	PostID uint `json:"post_id" form:"post_id"`
}

/**
 * @api {post} /api/v1/posttag Add posttag
 * @apiVersion 1.0.0
 * @apiName addPosttag
 * @apiGroup Posttag
 *
 * @apiParam {number} tag_id tag id
 * @apiParam {number} post_id post id
 *
 * @apiSuccess {String} message success message
 *
 * @apiError {String} error error message
 */

func addPosttag(c echo.Context) error {
	formData := new(posttagForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &posttag.Posttag{
		TagID:  formData.TagID,
		PostID: formData.PostID,
	}
	if err := p.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "posttag created successfully",
	})
}
