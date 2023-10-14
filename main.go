package main

import (
	"database/sql"
	"fmt"

	"github.com/abdelrhman-basyoni/golang-assignment/app/middlewares"
	Module_Book "github.com/abdelrhman-basyoni/golang-assignment/app/modules/book"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	const port = 3000
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware

	// Register the global error handler middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				middlewares.GlobalErrorHandler(err, c)
			}
			return nil
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//register  routes
	Module_Book.RegisterBookRoutes(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
