package handler

import (
	"goarch/database/rdb"
	"goarch/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	db := rdb.Local()
	u  := []model.User{}

	db.Find(&u)

	c.JSON(http.StatusOK, u)
}