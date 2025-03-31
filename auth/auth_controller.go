package auth

import (
	"fmt"
	"net/http"

	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	request := RequestRegisterUser{}
	ctx.BindJSON(&request)
	if err := request.validate(); err != nil {
		config.LoggerInited.Errorf("validation error: %v", err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := RegisterUserService(&request); err != nil {
		config.LoggerInited.Errorf("validation error: %v", err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	config.SendSuccess(ctx, "Register-User", fmt.Sprintf("user: %v created successfully", request.Email))
}

func LoginUser(ctx *gin.Context) {
	request := RequestLoginUser{}
	ctx.BindJSON(&request)
	if err := request.validate(); err != nil {
		config.LoggerInited.Errorf("validation error: %v", err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := LoginUserService(&request)
	if err != nil {
		config.LoggerInited.Errorf("login error: %v", err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Gerar JWT e retornar
	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		config.LoggerInited.Errorf("login error: %v", err.Error())
		config.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	config.SendSuccess(ctx, "Login-user", gin.H{
		"token": token,
	})
}
