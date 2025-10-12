package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sugito75/chat-app-server/config"
	"github.com/sugito75/chat-app-server/internal/user"

	mocks "github.com/sugito75/chat-app-server/pkg/mock"
)

func setupTestApp(service user.UserService) *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	handler := user.NewHandler(service)

	app.Post("/users", handler.CreateUser)
	app.Post("/login", handler.GetUserCredentials)

	app.Get("/check/:phone", handler.CheckIsNumberRegistered)

	app.Use(func(ctx *fiber.Ctx) error {
		return nil
	})

	return app
}

func TestHandleCreateUser(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		mockService := new(mocks.MockUserService)

		dto := user.CreateUserDTO{
			Username: "Alan",
			Password: "secret123",
			Phone:    "08123456789",
			Bio:      "Hi there!",
		}
		mockService.On("CreateUser", mock.AnythingOfType("user.CreateUserDTO")).Return(uint(1), nil)

		app := setupTestApp(mockService)

		body, _ := json.Marshal(dto)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, 201, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("should error when invalid body", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		app := setupTestApp(mockService)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString("invalid-json"))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should error when something wrong in service", func(t *testing.T) {
		mockService := new(mocks.MockUserService)

		mockService.On("CreateUser", mock.AnythingOfType("user.CreateUserDTO")).
			Return(uint(0), errors.New("db error"))

		app := setupTestApp(mockService)

		body, _ := json.Marshal(user.CreateUserDTO{
			Username: "Alan",
			Password: "secret123",
			Phone:    "08123456789",
		})
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandleGetUserCredentials(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		mockService := new(mocks.MockUserService)

		mockCred := &user.UserCredentialsDTO{
			ID:        1,
			Username:  "Alan",
			Phone:     "08123456789",
			SessionID: "session-abc",
		}
		mockService.On("GetUserCredentials", mock.AnythingOfType("user.GetUserCredentialsDTO")).
			Return(mockCred, nil)

		app := setupTestApp(mockService)

		body, _ := json.Marshal(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "secret123",
		})
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		app := setupTestApp(mockService)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString("invalid-json"))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 400 when credentials dont match", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		mockService.On("GetUserCredentials", mock.AnythingOfType("user.GetUserCredentialsDTO")).
			Return(nil, fiber.NewError(fiber.StatusBadRequest, "Invalid phone or password"))

		app := setupTestApp(mockService)

		body, _ := json.Marshal(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "wrongpass",
		})
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("should return 403  when user have no session id", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		mockService.On("GetUserCredentials", mock.AnythingOfType("user.GetUserCredentialsDTO")).
			Return(nil, fiber.NewError(fiber.StatusForbidden, "no session id"))

		app := setupTestApp(mockService)

		body, _ := json.Marshal(user.GetUserCredentialsDTO{
			Phone:    "08123456789",
			Password: "correctpass",
		})
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})
}

func TestHandleCheckIsNumberRegistered(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		app := setupTestApp(mockService)

		mockService.On("CheckIsNumberRegistered", "0813113").Return(true)

		req := httptest.NewRequest(http.MethodGet, "/check/0813113", nil)

		resp, _ := app.Test(req)

		assert.Equal(t, 200, resp.StatusCode)
	})
}
