package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/go-fiber-crud?charset=utf8mb4&parseTime=True&loc=Local"
	//const MONGODB = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	const DNS = MYSQL
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	fmt.Println("Connected to database.")
}
