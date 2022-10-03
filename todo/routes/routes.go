package Routes

import (
	Controllers "example/todolist/todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getalltodo", Controllers.GetTodos)
	r.GET("/gettodo_byuser/:user_id", Controllers.GetATodo)
	r.POST("/createtodo", Controllers.CreateATodo)

	r.PUT("/updatetodo/:user_id/:todo_id", Controllers.UpdateATodo)
	r.DELETE("/deletetodo/:todo_id", Controllers.DeleteATodo)
	return r
}
