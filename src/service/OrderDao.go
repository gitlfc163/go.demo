package service

import (
	"goormdemo1/dbutils"
	"goormdemo1/src/models"

	"gorm.io/gorm"
)

type OrderDao struct{}

func (dao *OrderDao) Create(model *models.Order) (tx *gorm.DB) {
	//通过数据的指针来创建
	result := dbutils.Db.Create(&model)
	return result
}

func (dao *OrderDao) Update(model models.Order) (tx *gorm.DB) {
	// Update - 更新多个字段 仅更新非零值字段
	result := dbutils.Db.Model(&model).Updates(model)
	return result
}

func (dao *OrderDao) Read(id uint) (model models.Order, tx *gorm.DB) {
	// 根据整型主键查找
	dbutils.Db.First(&model, id)
	return
}

// 根据主键删除
func (dao *OrderDao) Delete(id uint) (tx *gorm.DB) {
	var model models.Order
	// Delete - 删除 product
	dbutils.Db.Delete(&model, id)
	return
}
