package router

import (
	"goarch/handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *api) InitV1() {
	a.v1.GET("/users", handler.GetUserList)

	a.v1.GET("/status/:statuscode", func(c *gin.Context) {
		code, err := strconv.Atoi(c.Param("statuscode"))
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		if code < 100 || code > 600 {
			code = 200
		}

		c.JSON(code, gin.H{
			"result": "success",
		})
	})
}