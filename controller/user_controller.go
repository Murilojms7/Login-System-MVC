package controller

import (
	"fmt"
	"net/http"

	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/controller/request"
	"github.com/Murilojms7/LoginSystemMVC/service"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	users, err := service.AllUsers()
	if err != nil {
		config.LoggerInited.Error(err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var usersDTO []gin.H
	for _, user := range users {
		usersDTO = append(usersDTO, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email})
	}
	config.SendSuccess(ctx, "Get-Users", usersDTO)
}

func GetUserById(ctx *gin.Context) {
	requestId := ctx.Param("id")
	if requestId == "" {
		config.SendError(ctx, http.StatusBadRequest, "param: id, (type: queryParameter) is required")
		return
	}
	user, err := service.UserById(requestId)
	if err != nil {
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	config.SendSuccess(ctx, "show-user-by-id", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email})
}

func UpdateUserById(ctx *gin.Context) {
	requestUser := request.RequestUpdateUser{}
	ctx.BindJSON(&requestUser)
	if err := requestUser.Validate(); err != nil {
		config.GetLogger(fmt.Sprintf("validation error: %v", err.Error()))
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		config.GetLogger("Id not found")
		config.SendError(ctx, http.StatusBadRequest, "param: id, (type: queryParameter) is required")
		return
	}

	user, err := service.UpdateUserById(id, requestUser)
	if err != nil {
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	config.SendSuccess(ctx, "update-user", fmt.Sprintf("User: %v was updated", user.Name))
}

func DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		config.GetLogger("Id not found")
		config.SendError(ctx, http.StatusBadRequest, "param: id, (type: queryParameter) is required")
		return
	}
	if err := service.DeleteUser(id); err != nil {
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	config.SendSuccess(ctx, "delete-user", fmt.Sprintf("Delete user with id %v successfully", id))
}
