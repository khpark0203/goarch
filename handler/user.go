package handler

import (
	"goarch/message"
	"goarch/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	userSvc := service.NewUser()
	res, err := userSvc.GetList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	users := make([]message.User, len(res))
	for k, v := range res {
		users[k] = message.User{
			ID:        v.ID,
			Name:      v.Name,
			Age:       v.Age,
			CreatedAt: v.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, users)
}