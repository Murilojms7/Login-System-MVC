package router

import (
	"github.com/Murilojms7/LoginSystemMVC/auth"
	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/controller"
	"github.com/Murilojms7/LoginSystemMVC/server/middleware"
	"github.com/gin-gonic/gin"
)

func initializeUserRoutes(router *gin.Engine) {
	config.InitializeHandler()
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", auth.RegisterUser)
		authRouter.POST("/login", auth.LoginUser)
	}

	user := router.Group("/user", middleware.Auth())
	{
		user.GET("/:id", controller.GetUserById)
		user.GET("/", controller.GetAllUsers)

		user.PUT("/:id", controller.UpdateUserById)
		user.DELETE("/:id", controller.DeleteUser)
	}
}
