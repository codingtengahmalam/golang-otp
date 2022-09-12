package validator

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GoPlaygroundValidator struct {
	Validator *validator.Validate
}

func (gp *GoPlaygroundValidator) Validate(i interface{}) error {
	if err := gp.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
