package handler

import (
	"goarch/message"
	"goarch/service"
	"goarch/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.User
}

func NewUserHandler() userHandler {
	return userHandler{
		userService: *service.NewUser(),
	}
}

func (uh *userHandler) GetList(c *gin.Context) {
	res, msgerr := uh.userService.GetList()
	if msgerr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, msgerr)
		return
	}

	users := make([]message.UsersGetListRes, len(res))
	for k, v := range res {
		users[k] = message.UsersGetListRes{
			ID:        v.ID,
			Name:      v.Name,
			Age:       v.Age,
			CreatedAt: util.TimeToStringPtr(v.CreatedAt),
		}
	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) Create(c *gin.Context) {
}