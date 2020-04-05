package v1

import (
	"strconv"

	"github.com/MahdiRazaqi/nevees-backend/post"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * @api {get} /api/v1/post/:id Get a post
 * @apiVersion 1.0.0
 * @apiName onePost
 * @apiGroup Post
 *
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error error message
 */

func onePost(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(bson.M{"_id": id})
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"post": p,
	})
}

/**
 * @api {get} /api/v1/post List posts
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

	posts, err := post.Find(bson.M{}, page, limit)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"posts": posts,
	})
}
