package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"project/config"
	"project/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersController(t *testing.T) {
	// Setup
	config.InitDBTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err.Error())
		}

		assert.Equal(t, "success", response["message"])
		assert.NotNil(t, response["data"])
	}
}

func TestGetUserController(t *testing.T) {
	// Setup
	config.InitDBTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err.Error())
		}

		assert.Equal(t, "success get user by id", response["message"])
		assert.NotNil(t, response["user"])
	}
}

func TestCreateUserController(t *testing.T) {
	// Setup
	config.InitDBTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"name":"John Doe","email":"johndoe@example.com","password":"password"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err.Error())
		}

		assert.Equal(t, "success create user", response["message"])
		assert.NotNil(t, response["user"])
	}
}

func TestUpdateUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"name": "John Doe", "email": "johndoe@gmail.com", "password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"message":"success update user"`)
		assert.Contains(t, rec.Body.String(), `"name":"John Doe"`)
		assert.Contains(t, rec.Body.String(), `"email":"johndoe@gmail.com"`)
	}
}

func TestDeleteUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"message":"success delete user"`)
	}
}

func TestLoginUsersController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "johndoe@gmail.com", "password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	if assert.NoError(t, LoginUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"status":"success login"`)
		assert.Contains(t, rec.Body.String(), `"name":"John Doe"`)
		assert.Contains(t, rec.Body.String(), `"email":"johndoe@gmail.com"`)
		assert.Contains(t, rec.Body.String(), `"token":`)
	}
}
