package dbutils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB // 定义一个全局的DB，是一个连接池对象

//var db *sql.DB //指向数据库
const (
	server   = "localhost" // server
	port     = "3306"      // port
	user     = "root"      // user
	password = "123456"    // password
	database = "webdemo"   // database
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, server, port, database)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql error:", err)
		panic(err.Error())
	}
	fmt.Println("open mysql database is ok")
}
