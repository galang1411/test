package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
}
