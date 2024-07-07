package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100);unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}

type UserRegister struct {
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
