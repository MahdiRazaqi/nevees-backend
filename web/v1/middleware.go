package v1

import (
	"github.com/labstack/echo"
	"github.com/neveesco/nevees-backend/user"
)

func userRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &user.User{}
		if err := u.LoadByToken(user.GetToken(c.Request())); err != nil {
			return c.JSON(400, echo.Map{"error": "loading user from token " + err.Error()})
		}
		c.Set("user", u)
		return next(c)
	}
}
