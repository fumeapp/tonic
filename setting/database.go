package setting

type DatabaseSetting struct {
	Connect     string
	Host        string
	Port        string
	Database    string
	Username    string
	Password    string
	TablePrefix string
	Logging     string
}

var Database = &DatabaseSetting{}

func DatabaseSetup() *DatabaseSetting {

	Database.Connect = env("DB_CONNECT", "false")
	Database.Logging = env("DB_LOGGING", "false")
	Database.Host = env("DB_HOST", "localhost")
	Database.Port = env("DB_PORT", "3306")
	Database.Database = env("DB_DATABASE", "tonic")
	Database.Username = env("DB_USERNAME", "root")
	Database.Password = env("DB_PASSWORD", "")
	Database.TablePrefix = env("DB_PREFIX", "")

	return Database
}
