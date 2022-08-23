package service

import (
	"goormdemo1/dbutils"
	"goormdemo1/src/models"

	"gorm.io/gorm"
)

type ProductDao struct{}

func (dao *ProductDao) Create(product *models.Product) (tx *gorm.DB) {
	//通过数据的指针来创建
	result := dbutils.Db.Create(&product)
	return result
}

func (dao *ProductDao) Update(product models.Product) (tx *gorm.DB) {
	// Update - 更新多个字段 仅更新非零值字段
	result := dbutils.Db.Model(&product).Updates(product)
	return result
}

func (dao *ProductDao) Read(id uint) (product models.Product, tx *gorm.DB) {
	// 根据整型主键查找
	dbutils.Db.First(&product, id)
	return
}

//根据主键删除
func (dao *ProductDao) Delete(id uint) (tx *gorm.DB) {
	var product models.Product
	// Delete - 删除 product
	dbutils.Db.Delete(&product, id)
	return
}
