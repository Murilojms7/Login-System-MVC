package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	LoggerInited *Logger
)

func Init() error {
	if DB != nil && LoggerInited != nil {
		return nil
	}
	// Initialize Logger
	LoggerInited = NewLogger("main")

	// Initialize PostGre
	var err error
	DB, err = InitializePostgre()
	if err != nil {
		return fmt.Errorf("error initializing postgre: %v", err)
	}

	LoggerInited.Infof("Database initialized successfully")
	return nil
}

func GetPostgre() *gorm.DB {
	if DB == nil {
		panic("Database not initialized. Make sure to call config.Init() before using GetPostgre().")
	}
	return DB
}

func GetLogger(p string) *Logger {
	if LoggerInited == nil {
		LoggerInited = NewLogger(p)
	}
	return LoggerInited
}
