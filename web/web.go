package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	v1 "github.com/neveesco/nevees-backend/web/v1"
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
