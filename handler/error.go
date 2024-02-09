package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

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
