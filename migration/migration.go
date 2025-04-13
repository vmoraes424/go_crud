package main

import (
	"fmt"
	"go_crud/database"
	"go_crud/models"
)

func init() {
	database.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.Workout{})
	fmt.Println("Migration deu certo")
}
