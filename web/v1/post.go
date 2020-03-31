package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/post"
	"github.com/labstack/echo"
)

type postForm struct {
	Title   string   `json:"title" form:"title"`
	Content string   `json:"content" form:"content"`
	Tags    []string `json:"tags" form:"tags"`
}

func addPost(c echo.Context) error {
	formData := new(postForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		Title:   formData.Title,
		Content: formData.Content,
		Tags:    formData.Tags,
	}
	if err := p.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post created successfully",
		"post":    p,
	})
}
