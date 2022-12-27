package router

import "github.com/gin-gonic/gin"

type api struct {
	main *gin.Engine
	v1   *gin.RouterGroup
}

func InitGet() *api {
	// api router 시작 부분
	a := &api{
		main: gin.Default(),
	}
	a.v1 = a.main.Group("/api/v1")

	// 각 router 시작 부분
	a.InitV1()

	return a
}

func (a *api) Run(port string) {
	a.main.Run(port)
}