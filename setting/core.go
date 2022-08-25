package setting

type CoreSetting struct {
	WebURL string
}

var Core = &CoreSetting{}

func CoreSetup() *CoreSetting {
	Core.WebURL = env("WEB_URL", "http://localhost:3000")
	return Core
}
