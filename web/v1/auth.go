package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/user"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type authForm struct {
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
		Email:    formData.Email,
		Password: passwd.Make(formData.Password),
	}

	if err := u.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   "",
		"user":    u.Mini(),
	})
}

func login(c echo.Context) error {
	formData := new(authForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	u, err := user.LoadByEmail(formData.Email)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   "",
		"user":    u.Mini(),
	})
}
