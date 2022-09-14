package controllers

import (
	"github.com/ferriantitus/alterra-agmc/day2/static-api/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Book := models.Book{id, "Test", "ferrian"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": Book,
	})
}

func GetAllBooks(c echo.Context) error {
	Book := models.Book{
		ID:     1,
		Title:  "Test1",
		Author: "ferri",
	}
	return c.JSON(http.StatusOK, Book)
}

func DeleteBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully deleted",
	})
}

func CreateBooks(c echo.Context) error {
	Book := models.Book{
		ID:     2,
		Title:  "Test2",
		Author: "ferri",
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create book",
		"data": Book,
	})
}

func UpdateBook(c echo.Context) error {
	Book := models.Book{
		ID:     3,
		Title:  "Test3",
		Author: "ferri",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Edit book success",
		"data":    Book,
	})
}