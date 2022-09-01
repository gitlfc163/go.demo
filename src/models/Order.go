package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	LabId     int64
	UserID    int64
	ProductID int64
}
