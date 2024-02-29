package models

import "time"

type Product struct {
	ID  uint `json:"id" gorm:"primarykey"`
	createAt time.Time
	Name string `json:"name"`
	SerialNumber string `json:"serialNumber"`

}