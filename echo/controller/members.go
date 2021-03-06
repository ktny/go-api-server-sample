package controller

import (
	"net/http"

	"github.com/ktny/go-api-server-sample/echo/db"
	"github.com/labstack/echo"
)

func GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		members := []db.Member{}
		db := db.GetConnection()
		data := db.Find(&members)
		return c.JSON(http.StatusOK, data)
	}
}
