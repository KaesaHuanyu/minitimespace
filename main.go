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

	//APIs List Route
	// e.GET("/", apis)

	//Login Route
	e.GET("/login", h.Login)

	e.POST("/users", h.Protect(h.CreateUser))

	e.Start(":8823")
}
