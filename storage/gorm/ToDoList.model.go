package model

import (
	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	ID          string `gorm:"id, primaryKey"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Items       []ToDo `gorm:"foreignKey:ListId;references:ID"`
}
