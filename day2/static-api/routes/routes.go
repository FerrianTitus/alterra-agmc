package routes

import (
	"github.com/ferriantitus/alterra-agmc/day2/static-api/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	v1Books := e.Group("/v1")

	v1Books.GET("/books", controllers.GetAllBooks)
	v1Books.GET("/books/:id", controllers.GetBook)
	v1Books.DELETE("/books/:id", controllers.DeleteBooks)
	v1Books.PUT("/books/:id", controllers.UpdateBook)
	v1Books.POST("/books", controllers.CreateBooks)
	return e
}