package store

import (
	"github.com/jinzhu/gorm"
)

//Car Машинки
type Car struct {
	gorm.Model
	Name   string `json:"name"`
	Brand  string `json:"color"`
	Engine string `json:"engine"`
	Price  int    `json:"price"`
}

//Cars Много машинок
type Cars []Car
