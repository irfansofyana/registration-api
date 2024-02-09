package handler

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (s *Server) GetUserProfile(c echo.Context) error {
	headers := c.Request().Header

	token := headers.Get("Authorization")
	userId, err := validateAuthHeader(c, token)
	if err != nil {
		return newError(c, http.StatusUnauthorized, "invalid authorization token")
	}

	profile, err := s.Repository.GetProfile(c.Request().Context(), repository.GetProfileInput{
		Id: userId,
	})
	if err != nil {
		return newInternalServerError(c, err)
	}
	if profile == nil {
		return newError(c, http.StatusNotFound, "User is not found")
	}

	response := generated.ProfileResponse{
		FullName:    profile.FullName,
		PhoneNumber: profile.PhoneNumber,
	}

	return c.JSON(http.StatusOK, response)
}

func validateAuthHeader(ctx echo.Context, token string) (string, error) {
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return "", newError(ctx, http.StatusUnauthorized, "Authorization header is missing")
	}

	removedPrefix := strings.TrimPrefix(token, "Bearer ")

	return validateJwtToken(removedPrefix)
}

func validateJwtToken(tokenString string) (string, error) {
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(config.Instance.PublicKey)
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Issuer, nil
	}

	return "", errors.New("invalid authorization token")
}
