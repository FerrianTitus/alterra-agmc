package routes

import (
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/status", controllers.ShowStatus)

	v1 := e.Group("/v1")

	v1.POST("/login", controllers.V1LoginUsersController)

	v1Users := v1.Group("/users")

	v1Users.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		AuthScheme: "Bearer",
		Skipper: func(c echo.Context) bool {
			// Skip middleware if method is equal 'POST'
			if c.Request().Method == "POST" {
				return true
			}
			return false
		},
	}))

	v1Users.GET("", controllers.V1GetUsersController)
	v1Users.GET("/:id", controllers.V1GetUserByIdController)
	v1Users.POST("", controllers.V1CreateUserController)
	v1Users.DELETE("/:id", controllers.V1DeleteUserController)
	v1Users.PUT("/:id", controllers.V1UpdateUserByIdController)

	return e
}
