package main

import (
	Config "example/todolist/todo/config"
	Routes "example/todolist/todo/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	//run the migrations: todo struct
	//Config.DB.AutoMigrate(&Models.Tb_Todo{})

	//setup routes
	r := Routes.SetupRouter()

	//running
	r.Run()
}
