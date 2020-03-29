package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/user"
	"github.com/labstack/echo"
)

type authForm struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required"`
}

func register(c echo.Context) error {
	formData := new(authForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	u := &user.User{
		Username: formData.Username,
		Email:    formData.Email,
		Password: formData.Password,
	}

	if err := u.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   "",
		"user":    u,
	})
}
