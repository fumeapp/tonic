package setting

type CoreSetting struct {
	URL    string
	WebURL string
	ApiURL string
	AppENV string
}

var Core = &CoreSetting{}

func CoreSetup() *CoreSetting {
	Core.URL = env("URL", "http://localhost:3000")
	Core.WebURL = env("WEB_URL", "http://localhost:3000")
	Core.ApiURL = env("API_URL", "http://localhost:8000")
	Core.AppENV = env("APP_ENV", "dev")
	return Core
}
