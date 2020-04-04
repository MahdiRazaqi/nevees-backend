package v1

import (
	"strconv"

	"github.com/MahdiRazaqi/nevees-backend/post"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
