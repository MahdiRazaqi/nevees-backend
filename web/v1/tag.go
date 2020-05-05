package v1

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/tag"
)

type tagForm struct {
	Name string `json:"name" form:"name"`
}

/**
 * @api {post} /api/v1/tag Add tag
 * @apiVersion 1.0.0
 * @apiName addTag
 * @apiGroup Tag
 *
 * @apiParam {String} title tag title
 *
 * @apiSuccess {String} message success message
 * @apiSuccess {Object} tag tag model
 *
 * @apiError {String} error error message
 */

func addTag(c echo.Context) error {
	formData := new(tagForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t := &tag.Tag{
		Name: formData.Name,
	}
	if err := t.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "tag created successfully",
		"tag":     t,
	})
}

/**
 * @api {get} /api/v1/public/tag List tags
 * @apiVersion 1.0.0
 * @apiName listTags
 * @apiGroup Tag
 *
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 *
 * @apiSuccess {[]Object} tags array of tag model
 *
 * @apiError {String} error error message
 */

func listTags(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	tags, err := tag.Find(limit, page, "", "")
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"tags": tags,
	})
}
