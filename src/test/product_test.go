package test

import (
	"fmt"
	"goormdemo1/src/models"
	"goormdemo1/src/service"
	"testing"
)

// func TestMain(t *testing.M) {
// 	fmt.Println("start testing")
// 	initDB()
// }

// func TestUser(t *testing.T) {
// 	fmt.Println("start test user func")
// 	t.Run("test add user", TestAdduser)
// }

func TestAddProduct(t *testing.T) {
	fmt.Println("start test add Product:")
	product := models.Product{Code: "D42", Price: 100}
	dao := service.ProductDao{}
	//通过数据的指针来创建
	result := dao.Create(&product)
	if result != nil {
		fmt.Println("add Product: is error: ", result.Error)
	}
	fmt.Printf("add Product ID: %v\n", product.ID)
	fmt.Printf("add Product result: %v\n", result.RowsAffected)
	fmt.Println()
}
func TestReadProduct(t *testing.T) {
	fmt.Println("start test read Product:")
	dao := service.ProductDao{}
	product, result := dao.Read(2)
	if result != nil {
		fmt.Println("read Product: is error: ", result.Error)
	}
	fmt.Printf("read Product : %v\n", product)
	fmt.Println()
}
func TestUpdateProduct(t *testing.T) {
	fmt.Println("start test update Product:")
	dao := service.ProductDao{}
	product, result := dao.Read(2)
	if result != nil {
		fmt.Println("read Product: is error: ", result.Error)
		return
	}
	fmt.Printf("read Product : %v\n", product)

	product.Code = "D4C"
	product.Price = 106
	result = dao.Update(product)
	if result != nil {
		fmt.Println("update Product: is error: ", result.Error)
	}
	fmt.Printf("update Product result: %v\n", result.RowsAffected)
	fmt.Println()

}
func TestDeleteProduct(t *testing.T) {
	fmt.Println("start test delete Product:")
	dao := service.ProductDao{}
	result := dao.Delete(3)
	if result != nil {
		fmt.Println("delete Product: is error: ", result.Error)
	}
	fmt.Printf("delete Product result: %v", result)
	fmt.Println()
}
