package conndb

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func failErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Connection_db() *gorm.DB {
	dsn := "root:@tcp(localhost)/db_search?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	failErr(err)

	return db
}
