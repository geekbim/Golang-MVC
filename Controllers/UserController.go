package Controllers

import (
	"fmt"
	Models "mvc_pattern/Models/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get users
func GetUsers(c *gin.Context) {
	var user []Models.User

	err := Models.GetAllUsers(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Create user
func CreateUser(c *gin.Context) {
	var user Models.User

	c.BindJSON(&user)

	err := Models.CreateUser(&user)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Get user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")

	var user Models.User

	err := Models.GetUserByID(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Update user
func UpdateUser(c *gin.Context) {
	var user Models.User

	id := c.Params.ByName("id")

	err := Models.GetUserByID(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Delete user
func DeleteUser(c *gin.Context) {
	var user Models.User

	id := c.Params.ByName("id")

	err := Models.DeleteUser(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id " + id: "is deleted",
		})
	}
}
