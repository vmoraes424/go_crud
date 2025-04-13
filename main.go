package main

import (
	"go_crud/controllers"
	"go_crud/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnv()
	database.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/treinos", controllers.GetWorkouts)
	router.GET("/treino/:id", controllers.GetWorkout)
	router.DELETE("/treino/:id", controllers.DeleteWorkout)
	router.POST("/treino", controllers.PostWorkout)
	router.PUT("/treino/:id", controllers.EditWorkout)

	router.Run()
}
