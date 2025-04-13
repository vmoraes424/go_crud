package models

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	Name string
}
