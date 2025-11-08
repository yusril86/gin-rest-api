package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"column:title" json:"title"`
	Body  string `gorm:"column:body" json:"body"`
}
