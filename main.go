package main

import (
	_ "minitimespace/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	// h := handlers.New()
	// go h.AccessTokenUpdater()

	//APIs List Route
	// e.GET("/", apis)

	//Login Route
	// e.POST("/login", h.Login)

	//Unauthenticated Route

	//Restricted Route
	// r := e.Group("/restricted")
	// r.Use(middleware.JWT([]byte("secret")))
	// r.GET("", h.Restricted)

	e.Start(":8823")
}
