package config

import (
	"fmt"
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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