package v1

import (
	"github.com/MahdiRazaqi/nevees-backend/user"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type registerForm struct {
	Username string `json:"username" form:"username" validate:"required" `
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required"`
}

type loginForm struct {
	Username string `json:"username" form:"username" validate:"required" `
	Password string `json:"password" form:"password" validate:"required"`
}

func register(c echo.Context) error {
	formData := new(registerForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	u := &user.User{
		Username: formData.Username,
		Email:    formData.Email,
		Password: passwd.Make(formData.Password),
	}

	if err := u.Insert(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t, err := u.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   t,
		"user":    u.Mini(),
	})
}

func login(c echo.Context) error {
	formData := new(loginForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	u, err := user.AuthByUserPass(formData.Username, formData.Password)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t, err := u.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "login successfully",
		"token":   t,
		"user":    u.Mini(),
	})
}
