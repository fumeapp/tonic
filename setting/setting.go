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

func Setup() (
	*CoreSetting,
	*DatabaseSetting,
	*OpensearchSetting,
	*AWsSetting,
) {
	_ = godotenv.Load()
	return CoreSetup(), DatabaseSetup(), OpensearchSetup(), AwsSetup()
}

func IsDev() bool {
	return env("APP_ENV", "dev") == "dev"
}
func IsDebug() bool {
	return env("APP_DEBUG", "false") == "true"
}

func GetWebUrl() string {
	return env("WEB_URL", "http://localhost:3000")
}
