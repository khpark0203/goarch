package router

import "goarch/handler"

func (a *api) InitV1() {
	a.v1.GET("/users", handler.GetUserList)
}