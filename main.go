package main

import (
	"github.com/Murilojms7/LoginSystemMVC/config"
	"github.com/Murilojms7/LoginSystemMVC/router"
)

func main() {
	// initialize configs
	if err := config.Init(); err != nil {
		config.LoggerInited.Errorf("Config initialization error: %v", err)
		return
	}

	// initialize Router
	router.Initialize()
}
