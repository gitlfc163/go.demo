package test

import (
	"fmt"
	"goormdemo1/src/models"
	"goormdemo1/src/service"
	"testing"
)

func TestAddOrder(t *testing.T) {
	fmt.Println("start test add order:")
	model := models.Order{LabId: 1, UserID: 2, ProductID: 1}
	dao := service.OrderDao{}
	//通过数据的指针来创建
	result := dao.Create(&model)
	if result != nil {
		fmt.Println("add order: is error: ", result.Error)
	}
	fmt.Printf("add order ID: %v\n", model.ID)
	fmt.Printf("add order result: %v\n", result.RowsAffected)
	fmt.Println()

}
