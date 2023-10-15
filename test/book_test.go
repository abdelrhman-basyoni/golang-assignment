package Module_Book_test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	Module_Book "github.com/abdelrhman-basyoni/golang-assignment/app/modules/book"
	imp_repo "github.com/abdelrhman-basyoni/golang-assignment/core/implementation/repositories"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestController() (*sql.DB, *Module_Book.BookController) {
	db, err := sql.Open("sqlite3", "./test_database.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS books")
	if err != nil {
		panic(err)
	}

	ctrl := Module_Book.NewBookController(db)
	return db, ctrl
}

var db, ctrl = setupTestController()
var repo = imp_repo.NewBookSqlRepo(db)

func TestEndToEndCreateBook(t *testing.T) {

	// Create a new instance of the Echo application
	e := echo.New()

	// Create a request payload (JSON)
	payload := `{
        "Name": "End-to-End Test Book",
        "Genre": "Test Genre",
        "Price": 29.99
    }`

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Provide the database connection as needed
	err := ctrl.HandleCreate(c)

	// Assert the response and status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestHandleGetAll(t *testing.T) {

	repo := imp_repo.NewBookSqlRepo(db)
	// Create a new instance of the Echo application
	e := echo.New()
	// adding data to the data base
	repo.Create("End-to-End Test Book", "Test Genre", 29.99)
	// Define your expected response
	// expectedResponse := []domain_entities.Book{
	// 	// Define your expected book data here
	// 	{ID: "1", Name: "End-to-End Test Book", Genre: "Test Genre", Price: 29.99},
	// }

	// Perform the GET request to /books
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := ctrl.HandleGetAll(c)

	// Assert the response and status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Unmarshal the response and compare it to the expected response
	// var actualResponse []domain_entities.Book
	// err = json.NewDecoder(rec.Body).Decode(&actualResponse)
	// assert.NoError(t, err)
	// assert.Equal(t, expectedResponse, actualResponse)
}

func TestHandleGet(t *testing.T) {
	repo.Create("End-to-End Test Book 2", "Test Genre", 29.99)
	// Create a new instance of the Echo application
	e := echo.New()
	bookID := "2"
	// Define the book ID to retrieve

	// Perform the GET request to /books/get/:id
	req := httptest.NewRequest(http.MethodGet, "/books/"+bookID, nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues(bookID)
	err := ctrl.HandleGetById(c)

	// Assert the response and status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestHandleDelete(t *testing.T) {

	// Create a new instance of the Echo application
	e := echo.New()
	bookID := "1"
	// Perform the DELETE request to /books/delete/:id
	req := httptest.NewRequest(http.MethodDelete, "/books/"+bookID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues(bookID)
	err := ctrl.HandleDelete(c)

	// Assert the response and status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHandleUpdate(t *testing.T) {

	// Create a new instance of the Echo application
	e := echo.New()

	// Define the book ID to update
	bookID := "1" // Replace with a valid book ID

	// Define the update data as a JSON payload
	payload := `{
        "Name": "Updated Book Name",
        "Genre": "Updated Genre",
        "Price": 39.99
    }`

	// Perform the PUT request to /books/edit/:id
	req := httptest.NewRequest(http.MethodPut, "/books/"+bookID, bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues(bookID)
	err := ctrl.HandleUpdate(c)

	// Assert the response and status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}
