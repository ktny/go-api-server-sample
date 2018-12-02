package route

import (
	"github.com/ktny/go-api-server-sample/echo/controller"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/members", controller.GetMembers())
		// api.POST("/members", controller.PostMembers())
	}
	return e
}
