package Module_Book

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func RegisterBookRoutes(e *echo.Echo, db *sql.DB) {
	controller := NewBookController(db)
	bookGroup := e.Group("/books")
	bookGroup.POST("", controller.HandleCreate)
	bookGroup.GET("/:id", controller.HandleGetById)
	bookGroup.GET("", controller.HandleGetAll)
	bookGroup.PUT("/:id", controller.HandleUpdate)
	bookGroup.DELETE("/:id", controller.HandleDelete)
}
