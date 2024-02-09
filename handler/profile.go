package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) GetUserProfile(ctx echo.Context) error {
	profile, err := s.authCheck(ctx)
	if err != nil {
		return err.toEcho(ctx)
	}

	response := generated.ProfileResponse{
		FullName:    profile.FullName,
		PhoneNumber: profile.PhoneNumber,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *Server) PatchUserProfile(ctx echo.Context) error {
	existingProfile, err := s.authCheck(ctx)
	if err != nil {
		return err.toEcho(ctx)
	}

	var req generated.UpdateUserProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return newError(ctx, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	var updateRequest = repository.UpdateUserProfileInput{
		Id: existingProfile.Id,
	}
	if req.PhoneNumber != nil {
		existingUser, err := s.Repository.GetUserByPhoneNumber(
			ctx.Request().Context(), repository.GetUserByPhoneNumberInput{
				PhoneNumber: *req.PhoneNumber,
			})
		if err != nil {
			return newInternalServerError(ctx, err)
		}
		if existingUser != nil {
			return newError(ctx, http.StatusConflict, "Conflict. Phone number already registered")
		}
		updateRequest.PhoneNumber = *req.PhoneNumber
	}
	if req.FullName != nil {
		updateRequest.FullName = *req.FullName
	}

	if err := s.Repository.UpdateUserProfile(ctx.Request().Context(), updateRequest); err != nil {
		return newInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
