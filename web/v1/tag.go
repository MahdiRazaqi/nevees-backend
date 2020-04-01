package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/tag"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tagForm struct {
	Name string `json:"name" form:"name"`
}

func addTag(c echo.Context) error {
	formData := new(tagForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t := &tag.Tag{
		Name: formData.Name,
	}
	if err := t.InsertOne(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "tag created successfully",
		"tag":     t,
	})
}

func getOneTag(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t, err := tag.FindOne(bson.M{"_id": id})
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"post": t,
	})
}
