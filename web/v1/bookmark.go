package v1

import (
	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/bookmark"
	"github.com/neveesco/nevees-backend/user"
)

type bookmarkForm struct {
	PostID uint `json:"post_id" form:"post_id"`
}

/**
 * @api {post} /api/v1/bookmark Add to bookmark
 * @apiVersion 1.0.0
 * @apiName addToBookmark
 * @apiGroup Bookmark
 *
 * @apiParam {number} post_id post id
 *
 * @apiSuccess {String} message success message
 *
 * @apiError {String} error error message
 */

func addToBookmark(c echo.Context) error {
	u := c.Get("user").(*user.User)
	formData := new(bookmarkForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	b := &bookmark.Bookmark{
		PostID: formData.PostID,
		UserID: u.ID,
	}
	if err := b.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "add to bookmark successfully",
	})
}

/**
 * @api {delete} /api/v1/bookmark/:id Remove from bookmark
 * @apiVersion 1.0.0
 * @apiName removeFromBookmark
 * @apiGroup Bookmark
 *
 * @apiSuccess {String} message success message
 *
 * @apiError {String} error error message
 */

func removeFromBookmark(c echo.Context) error {
	u := c.Get("user").(*user.User)
	id := c.Param("id")

	if err := bookmark.Delete("id = ? and user_id = ?", id, u.ID); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "remove from bookmark successfully",
	})
}
