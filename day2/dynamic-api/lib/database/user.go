package database

import (
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/config"
	"github.com/ferriantitus/alterra-agmc/day2/dynamic-api/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func GetUsersByID(id string) (interface{}, error) {
	var users models.Users
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(id string) (interface{}, error) {
	var users models.Users
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUserByID(id string, email string, password string) (interface{}, error){
	var users models.Users
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(models.Users{Email: email, Password: password}).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(email string, password string) (interface{}, error) {

	var users = []models.Users{{Email: email, Password: password}}

	if err := config.DB.Create(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}