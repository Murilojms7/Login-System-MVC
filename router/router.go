package router

import (
	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	r := gin.Default()

	// Initialize Configs
	config.Init()

	// Initialize Routes
	initializeUserRoutes(r)

	r.Run(":8080")
}
