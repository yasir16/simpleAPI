package controller

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  len(users),
		}
	}

	c.JSON(http.StatusOK, result)

}
