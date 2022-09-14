package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ferriantitus/day-1/models"
)

var DB *gorm.DB

func InitDB(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		"root",
		"",
		"localhost",
		"3306",
		"agmc",
		"utf8mb4",
		"true",
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration(){
	DB.AutoMigrate(&models.Users{})
}