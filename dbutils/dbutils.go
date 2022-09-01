package dbutils

import (
	"fmt"
	"goormdemo1/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var Db *gorm.DB // 定义一个全局的DB，是一个连接池对象

// var db *sql.DB //指向数据库
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
	// 检查 `User` 对应的表是否存在
	// hasUser := Db.Migrator().HasTable(&models.User{})
	// if hasUser == false {
	// 	fmt.Println(" user table is not created")
	// }

	crreateOrderTb()

	//迁移
	Db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	//分表
	Db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "lab_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKCustom,
	}, "users"))

	Db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "lab_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKCustom,
	}, "products"))

	Db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "lab_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "orders"))

}

func crreateOrderTb() {
	for i := 0; i < 10; i += 1 {
		table := fmt.Sprintf("orders_%02d", i)
		//Db.Exec(`DROP TABLE IF EXISTS ` + table)
		Db.Exec(`CREATE TABLE ` + table + ` (
			id BIGSERIAL PRIMARY KEY,
			lab_id bigint,
			user_id bigint,
			product_id bigint
		)`)
	}
}
