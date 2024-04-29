package models

import (
	"github.com/jinzhu/gorm"
)

type CustomerWithId struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}
