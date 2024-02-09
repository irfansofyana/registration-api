package handler

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (s *Server) authCheck(ctx echo.Context) (*repository.GetProfileOutput, *customError) {
	headers := ctx.Request().Header

	token := headers.Get("Authorization")
	userId, err := validateAuthHeader(ctx, token)
	if err != nil {
		return nil, newCustomError(http.StatusForbidden, "Invalid authorization token: "+err.Error())
	}

	profile, err := s.Repository.GetProfile(ctx.Request().Context(), repository.GetProfileInput{
		Id: userId,
	})
	if err != nil {
		return nil, newCustomError(http.StatusInternalServerError, err.Error())
	}
	if profile == nil {
		return nil, newCustomError(http.StatusNotFound, "User not found")
	}

	return profile, nil
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
