package Controllers

import (
	Models "example/todolist/todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all todos
func GetAllUser(c *gin.Context) {
	var user []Models.Ms_user
	err := Models.GetAllUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(c *gin.Context) {
	user_id := c.Params.ByName("user_id")
	var user Models.Ms_user
	err := Models.GetAUser(&user, user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Create a User
func CreateUser(c *gin.Context) {
	var user Models.Ms_user
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Update an existing User
func UpdateUser(c *gin.Context) {
	var user Models.Ms_user
	user_id := c.Params.ByName("user_id")
	err := Models.GetAUser(&user, user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"Update User ID :" + user_id: "Success"})
}

// Delete a User
func DeleteUser(c *gin.Context) {
	var user Models.Ms_user
	user_id := c.Params.ByName("user_id")
	err := Models.DeleteUser(&user, user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"user_id :" + user_id: "deleted"})
	}
}
