package Models

import (
	Config "example/todolist/todo/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// fetch all todos at once
func GetAllTodo(todo *[]Tb_todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// get Todo By User
func GetATodo(todo *[]Tb_todo, user_id string) (err error) {

	if err := Config.DB.Where("user_id = ?", user_id).Find(todo).Error; err != nil {
		return err
	}
	// if err := Config.DB.Where("todo_id = ?", id).First(todo).Error; err != nil {
	// 	return err
	// }
	return nil
}

// get Todo By User & Todo
func GetTodo(todo *Tb_todo, user_id string, todo_id string) (err error) {

	if err := Config.DB.Where("user_id = ? AND todo_id = ?", user_id, todo_id).Find(todo).Error; err != nil {
		return err
	}
	// if err := Config.DB.Where("todo_id = ?", id).First(todo).Error; err != nil {
	// 	return err
	// }
	return nil
}

// insert a todo
func CreateATodo(todo *Tb_todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// update a todo
func UpdateATodo(todo *Tb_todo, user_id string, todo_id string) (err error) {

	fmt.Println(todo)

	Config.DB.Model(&todo).Where("user_id = ? AND todo_id = ?", user_id, todo_id).Updates(map[string]interface{}{"todo_title": todo.Todo_Title, "todo_desc": todo.Todo_Desc})

	//Config.DB.Save(todo)
	return nil
}

// delete a todo
func DeleteATodo(todo *Tb_todo, todo_id string) (err error) {
	Config.DB.Where("todo_id = ?", todo_id).Delete(todo)
	return nil
}
