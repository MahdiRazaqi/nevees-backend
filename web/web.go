package web

import (
	v1 "github.com/MahdiRazaqi/nevees-backend/web/v1"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Start server
func Start() {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}

	v1.Register(e)

	e.Start(":8080")
}
