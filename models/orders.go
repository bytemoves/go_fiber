package models

import "time"

type Order struct {
	ID  uint `json:"id" gorm:"primarykey"`
	createAt time.Time
	Productrefer int `json:"product_id"`
	Product Product `gorm:"foreignKey:Productrefer"`
	Userrefer int `json: "user_id"`
	User User `gorm:"foreignKey:Userrefer"`
}