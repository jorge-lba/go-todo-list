package model

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	ID          string `gorm:"id, primaryKey"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Done        bool   `gorm:"done; type:boolean; default:false"`
	ListId      string
}
