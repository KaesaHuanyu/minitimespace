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
	//e.GET("/users/:uid", h.Protect(h.GetUserDetail))

	//创建小时空
	e.POST("/timespaces", h.Protect(h.CreateTimespace))
	//更新小时空
	e.PUT("/timespaces/:tid", h.Protect(h.UpdateTimespace))
	//查询当前用户已加入的小时空
	e.GET("/timespaces", h.Protect(h.GetTimespace))
	//查看某小时空详情
	e.GET("/timespaces/:tid", h.Protect(h.GetTimespaceDetail))
	//删除小时空
	e.DELETE("/timespaces/:tid", h.Protect(h.DeleteTimespace))
	//检查当前用户是否已加入该小时空
	e.GET("/timespaces/:tid/user", h.Protect(h.WhetherCurrentUserJoined))
	//小时空添加当前用户
	e.PATCH("/timespaces/:tid/add", h.Protect(h.JoinTimespace))

	//e.GET("/labels", h.Protect(h.GetLabels))
	e.GET("/labels/:lid/timespace", h.Protect(h.GetTimespace))

	e.Start(":8823")
}
