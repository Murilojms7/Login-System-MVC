package auth

import (
	"errors"
	"fmt"

	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/model"
	"github.com/Murilojms7/LoginSystemMVC/utils"
	"gorm.io/gorm"
)

func RegisterUserService(user *RequestRegisterUser) error {
	var existingUser model.User
	result := config.GetPostgre().Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		config.LoggerInited.Errorf("Email already exists: %v", user.Email)
		return fmt.Errorf("email already exists")
	}

	password, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		config.LoggerInited.Errorf("Error hashing password")
		return err
	}

	newUser := &model.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: password,
	}

	if err := config.GetPostgre().Create(newUser).Error; err != nil {
		config.LoggerInited.Errorf("Error registering user")
		return err
	}

	config.LoggerInited.Infof("User registered successfully: %v", user.Email)
	return nil
}

func LoginUserService(request *RequestLoginUser) (model.User, error) {
	var user model.User
	result := config.GetPostgre().Where("email = ?", request.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return model.User{}, errors.New("email or password incorrect")
	}
	return user, nil
}
