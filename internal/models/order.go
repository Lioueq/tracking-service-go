package models

type Order struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"type:varchar(100)"`
	Status string `json:"status" gorm:"type:varchar(50)"`
	UserId int    `json:"userId" gorm:"type:int"`
}

type OrderCreate struct {
	Name string `json:"name"`
}

type OrderUpdate struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}
