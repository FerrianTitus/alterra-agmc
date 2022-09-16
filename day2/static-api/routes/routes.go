package routes

import (
	"github.com/ferriantitus/alterra-agmc/day2/static-api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func New() *echo.Echo {
	e := echo.New()

	v1Books := e.Group("/v1")

	v1Books.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		AuthScheme: "Bearer",
		Skipper: func(c echo.Context) bool {
			// Skip middleware if method is equal 'POST'
			if c.Request().Method == "GET" {
				return true
			}
			return false
		},
	}))

	v1Books.GET("/books", controllers.GetAllBooks)
	v1Books.GET("/books/:id", controllers.GetBook)
	v1Books.DELETE("/books/:id", controllers.DeleteBooks)
	v1Books.PUT("/books/:id", controllers.UpdateBook)
	v1Books.POST("/books", controllers.CreateBooks)
	return e
}