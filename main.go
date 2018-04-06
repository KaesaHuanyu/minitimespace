package main

import (
	"minitimespace/handlers"

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
	// e.POST("/login", h.Login)

	//Unauthenticated Route

	//Restricted Route
	// r := e.Group("/restricted")
	// r.Use(middleware.JWT([]byte("secret")))
	// r.GET("", h.Restricted)

	//User Route
	e.GET("/users/:userID", h.GetUser)
	e.GET("/users/:userID/teams", h.GetUserTeams)
	e.GET("/users/:userID/activities", h.GetUserActivities)
	// e.GET("/users/:userID/invitations", h.GetUserInvitations)
	e.GET("/users", h.GetUsers)
	e.POST("/users", h.CreateUser)
	e.PATCH("/users/:userID/joins/:teamID", h.UserJoinsTeam)
	e.PATCH("/users/:userID/exits/:teamID", h.UserExitsTeam)
	// e.PATCH("/users/:userID/accept/:activityID", h.UserAcceptsActivity)
	e.DELETE("/users/:userID", h.DeleteUser)

	//Team Route
	e.GET("/teams/:teamID", h.GetTeam)
	e.GET("/teams", h.GetTeams)
	e.POST("/teams", h.CreateTeam)
	e.PUT("/teams/:teamID", h.UpdateTeam)
	e.DELETE("/teams/:teamID", h.DeleteTeam)

	//Activity Route
	e.GET("/activities", h.GetAvtivities)
	e.GET("/activities/:activityID", h.GetAvtivity)
	e.POST("/activities", h.CreateActivity)
	e.PUT("/activities/:activityID", h.CreateActivity)
	e.DELETE("/activities/:activityID", h.DeleteActivity)

	e.Start(":8823")
}
