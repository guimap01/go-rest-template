package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Author    string `json:"author"`
	Name      string `json:"name"`
	PageCount int    `json:"pageCount"`
}
