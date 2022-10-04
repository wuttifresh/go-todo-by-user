package Models

import (
	Config "example/todolist/todo/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// fetch all todos at once
func GetAllUser(ms_user *[]Ms_user) (err error) {
	if err = Config.DB.Find(ms_user).Error; err != nil {
		return err
	}
	return nil
}

// get User by user_id
func GetAUser(ms_user *Ms_user, user_id string) (err error) {

	if err := Config.DB.Where("user_id = ?", user_id).Find(ms_user).Error; err != nil {
		return err
	}
	// if err := Config.DB.Where("todo_id = ?", id).First(todo).Error; err != nil {
	// 	return err
	// }
	return nil
}

// insert a user
func CreateUser(ms_user *Ms_user) (err error) {
	if err = Config.DB.Create(ms_user).Error; err != nil {
		return err
	}
	return nil
}

// update a todo
func UpdateUser(ms_user *Ms_user, user_id string) (err error) {

	fmt.Println(ms_user)

	Config.DB.Model(&ms_user).Where("user_id = ?", user_id).Updates(map[string]interface{}{"username": ms_user.Username, "userpass": ms_user.Userpass, "name": ms_user.Name, "surname": ms_user.Surname, "email": ms_user.Email, "tel": ms_user.Tel})

	//Config.DB.Save(todo)
	return nil
}

// delete a todo
func DeleteUser(ms_user *Ms_user, user_id string) (err error) {
	Config.DB.Where("user_id = ?", user_id).Delete(ms_user)
	return nil
}
