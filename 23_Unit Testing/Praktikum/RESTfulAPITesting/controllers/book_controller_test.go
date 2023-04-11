package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"project/config"
	"project/controllers"
	"project/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {
	// setup
	config.InitDBTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, controllers.GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookController(t *testing.T) {
	// setup
	config.InitDBTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// testing
	if assert.NoError(t, controllers.GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookController(t *testing.T) {
	// setup
	config.InitDBTest()
	e := echo.New()
	book := models.Books{
		Title:  "Test Book",
		Author: "Test Author",
	}
	bookJSON, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, controllers.CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success create new book", response["message"])

		// validate data inserted to DB
		var result models.Books
		config.DB.First(&result, response["book"].(map[string]interface{})["id"].(float64))
		assert.Equal(t, book.Title, result.Title)
		assert.Equal(t, book.Author, result.Author)
	}
}

func TestUpdateBookController(t *testing.T) {
	// setup
	config.InitDBTest()
	e := echo.New()
	book := models.Books{
		Title:  "Test Book",
		Author: "Test Author",
	}
	config.DB.Create(&book)
	book.Title = "Updated Test Book"
	bookJSON, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(book.ID), bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(book.ID))

	// testing
	if assert.NoError(t, controllers.UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success update book", response["message"])

		// validate data updated in DB
		var result models.Books
		config.DB.First(&result, book.ID)
		assert.Equal(t, book.Title, result.Title)
		assert.Equal(t, book.Author, result.Author)
	}
}

func TestDeleteBookController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create a mock database connection
	mockDB := &mocks.DB{}
	mockBook := models.Books{
		ID:     1,
		Title:  "Test Book",
		Author: "John Doe",
	}
	mockDB.On("First", &mockBook, 1).Return(nil)
	mockDB.On("Delete", &mockBook).Return(nil)

	// Inject the mock database connection
	config.DB = mockDB

	// Call the controller function
	if err := DeleteBookController(c); err != nil {
		t.Errorf("Error deleting book: %v", err)
	}

	// Check the response
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}
	expectedBody := `{"message":"success delete book"}`
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected response body %s but got %s", expectedBody, rec.Body.String())
	}

	// Verify the mock database connection was called
	mockDB.AssertExpectations(t)
}
