package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

func newError(c echo.Context, status int, message string) error {
	return c.JSON(status, generated.ErrorResponse{Message: message})
}

func newInternalServerError(c echo.Context, err error) error {
	return newError(c, http.StatusInternalServerError, "Internal Server Error: "+err.Error())
}
