package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func env(key string, config string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return config
}

func Setup() (*DatabaseSetting, *OpensearchSetting) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return DatabaseSetup(), OpensearchSetup()
}

func IsDev() bool {
	return env("APP_ENV", "dev") == "dev"
}
func IsDebug() bool {
	return env("APP_DEBUG", "false") == "true"
}
