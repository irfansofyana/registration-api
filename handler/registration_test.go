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

func TestRegistration(t *testing.T) {
	t.Run("Invalid payload", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.RegistrationRequest{
			FullName:    "Testing",
			Password:    "wrongpassword",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.RegisterUser(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Phone number already registered", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.RegistrationRequest{
			FullName:    "Testing full name",
			Password:    "Matematika2345!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			&repository.GetUserByPhoneNumberOutput{
				Id:       "random-id",
				FullName: "Testing full name",
				Password: "random-password",
			}, nil)

		err := server.RegisterUser(c)

		var errResponse generated.ErrorResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &errResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "Phone number already registered", errResponse.Message)
	})

	t.Run("Successfully register new user", func(t *testing.T) {
		e := echo.New()
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository.NewMockRepositoryInterface(controller)
		server := &Server{
			Repository: mockRepo,
		}

		body := generated.RegistrationRequest{
			FullName:    "Testing full name",
			Password:    "Matematika2345!",
			PhoneNumber: "+628111237878",
		}
		jsonBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewReader(jsonBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepo.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return(
			"random-id",
			nil,
		)
		mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(
			nil, nil)

		err := server.RegisterUser(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}
