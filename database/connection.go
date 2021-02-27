package database

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@/goauth"), &gorm.Config{})

	if err != nil {
		panic("cound not react the database")
	}

	DB = connection;

	connection.AutoMigrate(&models.User{});
	
}
