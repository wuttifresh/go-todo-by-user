package Routes

import (
	Controllers "example/todolist/todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	todo := r.Group("/todo")
	{
		todo.GET("/getalltodo", Controllers.GetTodos)
		todo.GET("/gettodo_byuser/:user_id", Controllers.GetATodo)
		todo.POST("/createtodo", Controllers.CreateATodo)
		todo.PUT("/updatetodo/:user_id/:todo_id", Controllers.UpdateATodo)
		todo.DELETE("/deletetodo/:todo_id", Controllers.DeleteATodo)
	}

	user := r.Group("/user")
	{
		user.GET("/getalluser", Controllers.GetAllUser)
		user.GET("/getuserbyid/:user_id", Controllers.GetUser)
		user.POST("/createuser", Controllers.CreateUser)
		user.PUT("/updateuser/:user_id", Controllers.UpdateUser)
		user.DELETE("/deleteuser/:user_id", Controllers.DeleteUser)
	}

	return r
}
