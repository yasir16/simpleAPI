package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasir16/simpleAPI/structs"
)

// GetUsers is for get list user
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

// GetUser is get single User
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// CreateUser is creating user
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
			"error":   err.Error(),
		})
	} else {
		idb.DB.Create(&user)
		result = gin.H{
			"result": user,
		}
		c.JSON(http.StatusOK, result)
	}

}

// UpdateUser is update single user by id
func (idb *InDB) UpdateUser(c *gin.Context) {
	var updateUser structs.User
	errParam := c.Bind(&updateUser)
	if errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	id := c.Param("id")
	var (
		user    structs.User
		newUser structs.User
		result  gin.H
	)
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newUser.FullName = updateUser.FullName
	newUser.Username = updateUser.Username
	newUser.Phone = updateUser.Phone
	newUser.Address = updateUser.Address
	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"data":   newUser,
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeleteUser is deleting single user by id
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&user).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
