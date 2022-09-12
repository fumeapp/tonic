package setting

import (
	"os"

	"github.com/joho/godotenv"
)

func env(key string, config string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return config
}

func Setup() (*CoreSetting, *DatabaseSetting, *OpensearchSetting, *AWsSetting) {

	if os.Getenv("_HANDLER") == "" {
		if err := godotenv.Load(); err != nil {
		}
	}
	return CoreSetup(), DatabaseSetup(), OpensearchSetup(), AwsSetup()
}

func IsDev() bool {
	return env("APP_ENV", "dev") == "dev"
}
func IsDebug() bool {
	return env("APP_DEBUG", "false") == "true"
}
