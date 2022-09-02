package service

import (
	"goormdemo1/dbutils"
	"goormdemo1/src/models"

	"gorm.io/gorm"
)

type ProductDao struct{}

func (dao *ProductDao) Create(model *models.Product) (tx *gorm.DB) {
	//通过数据的指针来创建
	result := dbutils.Db.Create(&model)
	return result
}

func (dao *ProductDao) Update(model models.Product) (tx *gorm.DB) {
	// Update - 更新多个字段 仅更新非零值字段
	result := GetDb(1).Updates(model)
	return result
}

func (dao *ProductDao) Read(id int64) (model models.Product, tx *gorm.DB) {
	// 根据整型主键查找
	dbutils.Db.First(&model, id)
	//dbutils.Db.Table("products_01").Where("lab_id", int64(1)).First(&model, id)
	return
}

// 根据主键删除
func (dao *ProductDao) Delete(id int64) (tx *gorm.DB) {
	var model models.Product
	// Delete - 删除 product
	GetDb(1).Delete(&model, id)
	return
}

func GetDb(lab_id int64) (tx *gorm.DB) {
	return dbutils.Db.Model(&models.Product{}).Where("lab_id", int64(lab_id))
}
