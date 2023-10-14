package Module_Book

import (
	"database/sql"
	"encoding/json"
	"net/http"

	domain_entities "github.com/abdelrhman-basyoni/golang-assignment/core/domain/entities"
	domain "github.com/abdelrhman-basyoni/golang-assignment/core/domain/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type BookController struct {
	uc *domain.BookUseCases
}

func NewBookController(db *sql.DB) *BookController {
	useCases := domain.NewBookUseCases(db)

	return &BookController{uc: useCases}
}

func (bc *BookController) HandleCreate(c echo.Context) error {

	var book domain_entities.Book

	// Bind and validate the request body
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	// use the validator library to validate required fields
	if err := validate.Struct(&book); err != nil {

		return err
	}

	if err := bc.uc.Create(book.Name, book.Genre, book.Price); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)

}

func (bc *BookController) HandleGetAll(c echo.Context) error {

	books := bc.uc.GetAll()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"books": books}, "")

}

func (bc *BookController) HandleGetById(c echo.Context) error {
	id := c.Param("id")

	book, err := bc.uc.GetByID(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)

	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"book": book}, "")

}

func (bc *BookController) HandleUpdate(c echo.Context) error {
	id := c.Param("id")
	var update map[string]interface{}

	// Read the request body and check if it's valid JSON
	if err := json.NewDecoder(c.Request().Body).Decode(&update); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body: " + err.Error(),
		})
	}

	err := bc.uc.Update(id, update)

	if err != nil {
		return err

	}
	return c.NoContent(http.StatusOK)
}

func (bc *BookController) HandleDelete(c echo.Context) error {
	id := c.Param("id")

	err := bc.uc.Delete(id)
	if err != nil {
		return err

	}
	return c.NoContent(http.StatusOK)

}
