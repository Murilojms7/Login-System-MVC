package service

import (
	"github.com/Murilojms7/LoginSystemMVC/controller/request"
	"github.com/Murilojms7/LoginSystemMVC/model"
	"github.com/Murilojms7/LoginSystemMVC/repository"
	"github.com/Murilojms7/LoginSystemMVC/utils"
)

func UserById(id string) (model.User, error) {
	user, err := repository.UserById(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func AllUsers() ([]model.User, error) {
	users, err := repository.AllUsers()
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func UpdateUserById(id string, requestUser request.RequestUpdateUser) (model.User, error) {
	user, err := repository.UserById(id)
	if err != nil {
		return model.User{}, err
	}
	if requestUser.Email != "" {
		user.Email = requestUser.Email
	}
	if requestUser.Name != "" {
		user.Name = requestUser.Name
	}
	if requestUser.Password != "" {
		user.Password, err = utils.GenerateHashPassword(requestUser.Password)
		if err != nil {
			return model.User{}, err
		}
	}

	if err := repository.UpdateUser(user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	user, err := repository.UserById(id)
	if err != nil {
		return err
	}

	if err := repository.DeleteUser(&user); err != nil {
		return err
	}
	return nil
}
