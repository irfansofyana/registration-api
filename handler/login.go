package handler

import (
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Server) LoginUser(ctx echo.Context) error {
	var req generated.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return newError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	user, err := s.Repository.GetUserByPhoneNumber(
		ctx.Request().Context(), repository.GetUserByPhoneNumberInput{
			PhoneNumber: req.PhoneNumber,
		})
	if err != nil {
		return newInternalServerError(ctx, err)
	}
	if user == nil {
		return newError(ctx, http.StatusNotFound, "User is not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return newError(ctx, http.StatusUnauthorized, "Invalid password")
	}

	token, err := createJwtToken(user.Id)
	if err != nil {
		return newInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, generated.LoginResponse{Token: token})
}

func createJwtToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.Instance.PrivateKey)
	if err != nil {
		return "", err
	}

	return token.SignedString(signKey)
}
