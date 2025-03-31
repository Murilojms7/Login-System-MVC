package config

func InitializeHandler() {
	LoggerInited = GetLogger("handler")
	DB = GetPostgre()
}
