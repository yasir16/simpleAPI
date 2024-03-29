package main

import (
	"fmt"
	"net/http"

	"github.com/yasir16/simpleAPI/config"
	"github.com/yasir16/simpleAPI/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setupServer().Run()
}
func setupServer() *gin.Engine {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.GET("/users", auth, inDB.GetUsers)
	router.GET("/user/:id", auth, inDB.GetUser)
	router.PUT("/user/:id", auth, inDB.UpdateUser)
	router.DELETE("/user/:id", auth, inDB.DeleteUser)
	router.POST("/user", auth, inDB.CreateUser)
	// router.Run(":3000")
	return router
}
func auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	// if token.Valid && err == nil {
	if token == "initestyasir" {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   "not authorized",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
