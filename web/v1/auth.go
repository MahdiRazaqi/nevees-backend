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

/**
 * @api {post} /api/v1/auth/register Register user
 * @apiVersion 1.0.0
 * @apiName register
 * @apiGroup User
 *
 * @apiParam {String} username unique username
 * @apiParam {String} email unique email
 * @apiParam {String} password password
 *
 * @apiSuccess {String} message success message
 * @apiSuccess {String} token user token access jwt
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error error message
 */

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
		Role:     "user",
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

/**
 * @api {post} /api/v1/auth/login Login user
 * @apiVersion 1.0.0
 * @apiName login
 * @apiGroup User
 *
 * @apiParam {String} username unique username
 * @apiParam {String} password password
 *
 * @apiSuccess {String} message success message
 * @apiSuccess {String} token user token access jwt
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error error message
 */

func login(c echo.Context) error {
	formData := new(loginForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	u := &user.User{}
	if err := u.AuthByUserPass(formData.Username, formData.Password); err != nil {
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
