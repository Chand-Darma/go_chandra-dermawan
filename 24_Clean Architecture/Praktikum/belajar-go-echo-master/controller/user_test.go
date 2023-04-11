package controller_test

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/middleware"
	"belajar-go-echo/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"luasTrapesium/go/pkg/mod/gorm.io/driver/mysql@v1.4.7"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllUsers_Success(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&model.User{})
	db.Create(&model.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})

	// Setup request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute function
	handler := controller.GetAllUsers(db)
	err = handler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	expectedBody := `{"data":[{"id":1,"name":"John Doe","email":"john.doe@example.com","token":""}]}`
	assert.Equal(t, expectedBody, rec.Body.String())
}

func TestGetAllUsers_Error(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup database with error
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&model.User{})
	expectedErr := errors.New("test error")
	db.Find(&model.User{})
	db.Error = expectedErr

	// Setup request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute function
	handler := controller.GetAllUsers(db)
	err = handler(c)

	// Assertions
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"test error"}`
	assert.Equal(t, expectedBody, rec.Body.String())
}

func TestCreateUser(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john@example.com","password":"password"}`))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// create mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %s", err)
	}
	defer mockDB.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %s", err)
	}

	// create mock token
	token := "test_token"
	middleware.CreateToken = func(id int, email string) (string, error) {
		return token, nil
	}

	// expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// test CreateUser
	handler := CreateUser(db)
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response body
		expected := `{"data":{"id":1,"email":"john@example.com","token":"test_token"}}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
