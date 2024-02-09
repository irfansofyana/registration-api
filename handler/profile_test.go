package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProfile(t *testing.T) {
	t.Run("No authorization header", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.GetUserProfile(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})

	t.Run("Invalid authorization token", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer invalid-token")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.GetUserProfile(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})

	t.Run("User not found", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		token := getTokenForTest(t)
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(
			nil, nil)

		err := server.GetUserProfile(c)

		var errResponse generated.ErrorResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &errResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "User is not found", errResponse.Message)
	})

	t.Run("Successfully", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		token := getTokenForTest(t)
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(
			&repository.GetProfileOutput{
				FullName:    "Testing full name",
				PhoneNumber: "+628111237878",
			}, nil)

		err := server.GetUserProfile(c)

		var profileResponse generated.ProfileResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &profileResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Testing full name", profileResponse.FullName)
		assert.Equal(t, "+628111237878", profileResponse.PhoneNumber)
	})
}

func getTokenForTest(t *testing.T) string {
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

	_ = server.LoginUser(c)

	var loginResponse generated.LoginResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &loginResponse)

	return loginResponse.Token
}
