package Controllers

import (
	Models "example/todolist/todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all todos
func GetTodos(c *gin.Context) {
	var todo []Models.Tb_todo
	err := Models.GetAllTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	user_id := c.Params.ByName("user_id")
	var todo []Models.Tb_todo
	err := Models.GetATodo(&todo, user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Create a Todo
func CreateATodo(c *gin.Context) {
	var todo Models.Tb_todo
	c.BindJSON(&todo)
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func UpdateATodo(c *gin.Context) {
	var todo Models.Tb_todo
	todo_id := c.Params.ByName("todo_id")
	user_id := c.Params.ByName("user_id")
	err := Models.GetTodo(&todo, user_id, todo_id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, user_id, todo_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"Update Todo ID :" + todo_id: "Success"})
}

// Delete a Todo
func DeleteATodo(c *gin.Context) {
	var todo Models.Tb_todo
	todo_id := c.Params.ByName("todo_id")
	err := Models.DeleteATodo(&todo, todo_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"todo_id :" + todo_id: "deleted"})
	}
}
