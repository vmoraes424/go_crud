package controllers

import (
	"go_crud/database"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

func GetWorkouts(c *gin.Context) {
	var treinos []models.Workout

	result := database.DB.Find(&treinos)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"treinos": treinos,
	})
}

func PostWorkout(c *gin.Context) {

	var body struct {
		Name string
	}

	c.Bind(&body)

	treino := models.Workout{Name: body.Name}

	result := database.DB.Create(&treino)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"treino": treino,
	})
}

func GetWorkout(c *gin.Context) {

	id := c.Param("id")

	var treino models.Workout
	result := database.DB.First(&treino, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"treino": treino,
	})
}

func DeleteWorkout(c *gin.Context) {
	var treino models.Workout

	id := c.Param("id")

	result := database.DB.Delete(&treino, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"treino": treino,
	})
}

func EditWorkout(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name string
	}

	c.Bind(&body)

	var treino models.Workout
	database.DB.First(&treino, id)

	database.DB.Model(&treino).Updates(models.Workout{Name: body.Name})

	c.JSON(200, gin.H{
		"treino": treino,
	})
}
