package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("Internal server error", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.LoginRequest{
			Password:    "Matematika2345!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			nil, errors.New("some error"))

		err := server.LoginUser(c)

		var errResponse generated.ErrorResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &errResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Something unexpected happened. We're investigating the issue.", errResponse.Message)
	})

	t.Run("User not found", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.LoginRequest{
			Password:    "Matematika2345!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			nil, nil)

		err := server.LoginUser(c)

		var errResponse generated.ErrorResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &errResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "User is not found", errResponse.Message)
	})

	t.Run("Invalid password", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.LoginRequest{
			Password:    "WrongPassword123!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			&repository.GetUserByPhoneNumberOutput{
				Id:       "random-id",
				FullName: "Testing full name",
				Password: "$2a$10$.URirH.CXSA6CBnxi3d7Ceo7GB39Q7F9Uw/bJ7lhUQpnMjDHUuy.a",
			}, nil)

		err := server.LoginUser(c)

		var errResponse generated.ErrorResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &errResponse)

		assert.NoError(t, err)
		assert.Equal(t, "Invalid password", errResponse.Message)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Successfully login", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.LoginRequest{
			Password:    "Testing123!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			&repository.GetUserByPhoneNumberOutput{
				Id:       "random-id",
				FullName: "Testing full name",
				Password: "$2a$10$.URirH.CXSA6CBnxi3d7Ceo7GB39Q7F9Uw/bJ7lhUQpnMjDHUuy.a",
			}, nil)
		mockRepo.EXPECT().UpdateUserLoginCount(gomock.Any(), gomock.Any()).Return(nil)

		err := server.LoginUser(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
