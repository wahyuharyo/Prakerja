package config

import (
	"github.com/wahyuharyo/prakerja/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB 

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:0852@tcp(localhost:3306)/prakerja"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = database
	
	DB.AutoMigrate(&models.Menu{})
}