package repository

import (
	"errors"
	"fmt"

	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AllUsers() ([]model.User, error) {
	users := []model.User{}
	if err := config.GetPostgre().Find(&users).Error; err != nil {
		config.LoggerInited.Errorf("error searching user")
		return []model.User{}, fmt.Errorf("error searching user")
	}
	return users, nil
}

func UserById(id string) (model.User, error) {
	user := model.User{}
	if err := config.GetPostgre().First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.LoggerInited.Errorf("user not found")
			return model.User{}, fmt.Errorf("user not found")
		}
	}
	if user.ID == uuid.Nil {
		config.LoggerInited.Errorf("user not found")
		return model.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func UpdateUser(user model.User) error {
	if err := config.GetPostgre().Save(&user).Error; err != nil {
		config.LoggerInited.Errorf("error updating user")
		return fmt.Errorf("error updating user")
	}
	if user.ID == uuid.Nil {
		config.LoggerInited.Errorf("user not found")
		return fmt.Errorf("user not found")
	}
	return nil
}

func DeleteUser(user *model.User) error {
	if err := config.GetPostgre().Delete(&user).Error; err != nil {
		config.LoggerInited.Errorf("error deleting user")
		return fmt.Errorf("error deleting user")
	}
	return nil
}
