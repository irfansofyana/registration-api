package handler

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) RegisterUser(ctx echo.Context) error {
	var req generated.RegisterUserJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return newError(ctx, http.StatusBadRequest, "Invalid request")
	}

	if err := validateRegistrationRequest(req); err != nil {
		return newError(ctx, http.StatusBadRequest, err.Error())
	}

	existingUser, err := s.Repository.GetUserByPhoneNumber(
		ctx.Request().Context(),
		repository.GetUserByPhoneNumberInput{
			PhoneNumber: req.PhoneNumber,
		},
	)
	if err != nil {
		return newInternalServerError(ctx, err)
	}

	if existingUser != nil {
		return newError(ctx, http.StatusBadRequest, "Phone number already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return newInternalServerError(ctx, err)
	}

	user := repository.SaveUserInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashedPassword),
	}
	id, err := s.Repository.SaveUser(ctx.Request().Context(), user)
	if err != nil {
		return newInternalServerError(ctx, err)
	}

	var resp = generated.RegistrationResponse{
		UserId: id,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func validateRegistrationRequest(req generated.RegisterUserJSONRequestBody) error {
	if !validatePassword(req.Password) {
		return errors.New("password must contain at least one uppercase letter, one digit, and one special character")
	}
	return nil
}

func validatePassword(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`\W`).MatchString(password)

	return hasUpper && hasDigit && hasSpecial
}
