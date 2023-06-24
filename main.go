package main

import (
	"log"
	"net/http"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	adapter, err := gormadapter.NewAdapterByDB(DB)
	if err != nil {
		log.Fatalln("not able to connect")
	}
	e.Use(Authenticate(adapter))

	e.GET("/project", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "project get allowed")
	})
	e.POST("/project", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "project post allowed")
	})

	e.GET("/channel", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "channel get allowed")
	})

	e.POST("/channel", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "channel post allowed")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
