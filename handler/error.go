package handler

import (
	"errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type customError struct {
	code    int
	message string
}

func newCustomError(code int, message string) *customError {
	return &customError{code: code, message: message}
}

func (cErr *customError) toEcho(c echo.Context) error {
	if cErr.code == http.StatusInternalServerError {
		return newInternalServerError(c, errors.New(cErr.message))
	}

	return newError(c, cErr.code, cErr.message)
}

func newError(c echo.Context, status int, message string) error {
	return c.JSON(status, generated.ErrorResponse{Message: message})
}

func newInternalServerError(c echo.Context, err error) error {
	log.Errorf("Internal server error: %v", err)
	return newError(c,
		http.StatusInternalServerError,
		"Something unexpected happened. We're investigating the issue.",
	)
}
