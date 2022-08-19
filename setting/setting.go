package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Database = &DatabaseSetting{}

func env(key string, config string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return config
}

func Setup() *DatabaseSetting {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return DatabaseSetup()
}

func IsDev() bool {
	return env("APP_ENV", "dev") == "dev"
}
func IsDebug() bool {
	return env("APP_DEBUG", "false") == "true"
}
