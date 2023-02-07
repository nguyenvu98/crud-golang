package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:nguyenvu98@tcp(localhost:3306)/golang_crud"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Album{})

	DB = database
}
