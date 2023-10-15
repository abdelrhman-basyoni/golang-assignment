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
	e.HTTPErrorHandler = middlewares.GlobalErrorHandler
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//register  routes
	Module_Book.RegisterBookRoutes(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
