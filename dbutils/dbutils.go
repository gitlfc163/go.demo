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
	//按lab_id自动创建orders的相关折分表
	CreateOrderTable()
	CreateProductTable()

	//自动迁移
	Db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	//分表规则
	ShardingTables()
}

// 按lab_id自动创建orders的相关折分表
func CreateOrderTable() {
	for i := 1; i < 11; i += 1 {
		tableName := fmt.Sprintf("orders_%02d", i)
		hasUser := Db.Migrator().HasTable(tableName)
		if hasUser {
			fmt.Println(tableName, " table is created")
			continue
		}
		fmt.Println("create table", tableName)
		//Db.Exec(`DROP TABLE IF EXISTS ` + table)
		Db.Exec(`CREATE TABLE ` + tableName + ` (
			id bigint UNSIGNED NOT NULL AUTO_INCREMENT,
			lab_id bigint,
			user_id bigint,
			product_id bigint,
			created_at datetime(3) NULL DEFAULT NULL,
			updated_at datetime(3) NULL DEFAULT NULL,
			deleted_at datetime(3) NULL DEFAULT NULL,
			PRIMARY KEY (id) USING BTREE
		)`)
	}
}

// 按lab_id自动创建products的相关折分表
func CreateProductTable() {
	for i := 1; i < 11; i += 1 {
		tableName := fmt.Sprintf("products_%02d", i)
		hasUser := Db.Migrator().HasTable(tableName)
		if hasUser {
			fmt.Println(tableName, " table is created")
			continue
		}
		fmt.Println("create table", tableName)
		//Db.Exec(`DROP TABLE IF EXISTS ` + table)
		Db.Exec(`CREATE TABLE ` + tableName + ` (
			id bigint UNSIGNED NOT NULL AUTO_INCREMENT,
			lab_id bigint,
			e_name varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
			c_name varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
			code varchar(50),
			price float DEFAULT NULL,
			created_at datetime(3) NULL DEFAULT NULL,
			updated_at datetime(3) NULL DEFAULT NULL,
			deleted_at datetime(3) NULL DEFAULT NULL,
			PRIMARY KEY (id) USING BTREE
		)`)
	}
}

// 分表规则
func ShardingTables() {

	Db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "lab_id", //指定要分片的表字段名称
		NumberOfShards:      10,       //指定要分片的表数量
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "products"))

	Db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "lab_id",
		NumberOfShards:      10, //指定要分片的表数量，就是获取租户信息表的所有的lab_id进行创建相关折分的表
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "orders"))

	// Db.Use(sharding.Register(sharding.Config{
	// 	ShardingKey:         "lab_id",
	// 	NumberOfShards:      20,//指定要分片的表数量
	// 	PrimaryKeyGenerator: sharding.PKCustom,
	// }, "users"))
}
