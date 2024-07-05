package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status string `json:"string"`
}
