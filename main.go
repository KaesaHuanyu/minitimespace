package main

import (
	"minitimespace/handlers"
	_ "minitimespace/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	h := handlers.New()
	go h.AccessTokenUpdater()

	//Login Route
	e.GET("/login", h.Login)

	e.POST("/users", h.Protect(h.CreateUser))
	//timespace add me
	e.GET("/users/:uid", h.Protect(h.GetUserDetail))

	e.GET("/timespace", h.Protect(h.GetTimespace))
	e.POST("/timespace", h.Protect(h.CreateTimespace))
	e.DELETE("/timespace/:tid", h.Protect(h.DeleteTimespace))
	e.PUT("/timespace/:tid", h.Protect(h.UpdateTimespace))
	e.PATCH("/timespace/:tid/add/")
	e.GET("/timespace/:tid", h.Protect(h.GetTimespaceDetail))

	e.GET("/labels", h.Protect(h.GetLabels))
	e.GET("/labels/:lid/timespace", h.Protect(h.GetTimespace))

	e.Start(":8823")
}
