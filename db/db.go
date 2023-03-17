package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=1234 dbname=gorm  port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Fail to connect db. Error : " + err.Error())
		return
	}
	DB = db
	fmt.Println("Connect database successfully!")
}

func GetDB() *gorm.DB {
	if DB == nil {
		fmt.Println("DB nil")
	}
	return DB
}
