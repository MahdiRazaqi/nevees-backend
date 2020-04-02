package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/post"
	"github.com/MahdiRazaqi/nevees-backend/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postForm struct {
	Title   string   `json:"title" form:"title"`
	Content string   `json:"content" form:"content"`
	Tags    []string `json:"tags" form:"tags"`
}

func addPost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	formData := new(postForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		Title:   formData.Title,
		Content: formData.Content,
		User:    u.ID,
		Tags:    formData.Tags,
	}
	if err := p.InsertOne(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post created successfully",
		"post":    p,
	})
}

func removePost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := post.DeleteOne(bson.M{"_id": id, "_user": u.ID}); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post removed successfully",
	})
}

func editPost(c echo.Context) error {
	u := c.Get("user").(*user.User)
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	formData := new(postForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	filter := bson.M{"_id": id, "_user": u.ID}

	p, err := post.FindOne(filter)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p.Title = formData.Title
	p.Content = formData.Content
	p.Tags = formData.Tags

	if err := p.UpdateOne(filter); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post updated successfully",
		"post":    p,
	})
}

func getOnePost(c echo.Context) error {
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

func getAllPosts(c echo.Context) error {
	posts, err := post.Find(bson.M{})
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"posts": posts,
	})
}
