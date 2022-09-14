package controllers

import (
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/lib/database"
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/models"
	"github.com/labstack/echo"
	"net/http"
)

func V1GetUserByIdController(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	//for _, users := range users {
	//	if users.ID == (id) {
	//		return c.JSON(http.StatusOK, map[string]interface{}{
	//			"message": "success",
	//			"code": http.StatusOK,
	//			"data": users,
	//		})
	//	}
	//}
	//return c.JSON(http.StatusBadRequest, map[string]interface{}{
	//	"message": "user not found",
	//	"code": http.StatusBadRequest,
	//})

	id := c.Param("id")

	users, err := database.GetUsersByID(id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"code": http.StatusNotFound,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": users,
	})
}

func V1GetUsersController(c echo.Context) error {
	//email := c.QueryParam("email")
	//if email == ""{
	//	email = c.FormValue("email")
	//}
	//if email != "" {
	//	for _, users := range users {
	//		if users.Email == email {
	//			return c.JSON(http.StatusOK, map[string]interface{}{
	//				"message": "success",
	//				"code": http.StatusOK,
	//				"data": users,
	//			})
	//		}
	//	}
	//}
	users, err := database.GetUsers()
	if err != nil{
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": users,
	})
}

func V1CreateUserController(c echo.Context)error{
	users := models.Users{}
	c.Bind(&users)

	user, err := database.CreateUser(users.Email, users.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create user",
		"data": user,
	})
}

func V1DeleteUserController(c echo.Context)error {
	id := c.Param("id")
	_, err := database.DeleteUser(id)
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"code": http.StatusNotFound,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully deleted",
	})
}

func V1UpdateUserByIdController(c echo.Context) error {
	user := models.Users{}
	id := c.Param("id")

	c.Bind(&user)

	users, err := database.UpdateUserByID(id, user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Edit users success",
		"data":    users,
	})
}