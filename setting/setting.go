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

func Setup(file string) (*CoreSetting, *DatabaseSetting, *OpensearchSetting, *AWsSetting) {

	if err := godotenv.Load(file); err != nil {
		return CoreSetup(), DatabaseSetup(), OpensearchSetup(), AwsSetup()
	}
	return CoreSetup(), DatabaseSetup(), OpensearchSetup(), AwsSetup()
}

func IsDev() bool {
	return env("APP_ENV", "dev") == "dev"
}
func IsDebug() bool {
	return env("APP_DEBUG", "false") == "true"
}

func IsStaging() bool {
	return env("APP_ENV", "dev") == "staging"
}

func IsProduction() bool {
	return env("APP_ENV", "dev") == "production"
}

func IsTesting() bool {
	return env("APP_ENV", "dev") == "testing"
}
