package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type member struct {
	ID   int    `json:id`
	Name string `json:name`
}

func GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		members := []member{
			{ID: 1, Name: "ユーザー1"},
			{ID: 2, Name: "ユーザー2"},
		}
		return c.JSON(http.StatusOK, members)
	}
}
