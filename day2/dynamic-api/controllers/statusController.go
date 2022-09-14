package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func ShowStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}