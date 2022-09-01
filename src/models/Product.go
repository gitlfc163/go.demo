package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	LabId int64
	CName string
	EName string
	Code  string
	Price float32
}
