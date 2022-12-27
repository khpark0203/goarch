package handler

import (
	"goarch/message"
	"goarch/service"
	"goarch/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	userSvc := service.NewUser()
	
	req := message.GetUserListReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		msgerr := message.Error(message.ERROR_COMMON_INTERNAL, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, msgerr)
		return
	}
	
	res, msgerr := userSvc.GetList()
	if msgerr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, msgerr)
		return
	}

	users := make([]message.GetUserListRes, len(res))
	for k, v := range res {
		users[k] = message.GetUserListRes{
			ID:        v.ID,
			Name:      v.Name,
			Age:       v.Age,
			CreatedAt: util.TimeToStringPtr(v.CreatedAt),
		}
	}

	c.JSON(http.StatusOK, users)
}