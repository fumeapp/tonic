package setting

type DatabaseSetting struct {
	Driver      string
	Host        string
	Database    string
	Username    string
	Password    string
	TablePrefix string
	Logging     string
}

var Database = &DatabaseSetting{}

func DatabaseSetup() *DatabaseSetting {

	Database.Logging = env("DB_LOGGING", "false")
	Database.Driver = env("DB_DRIVER", "mysql")
	Database.Host = env("DB_HOST", "localhost")
	Database.Database = env("DB_DATABASE", "tonic")
	Database.Username = env("DB_USERNAME", "root")
	Database.Password = env("DB_PASSWORD", "")
	Database.TablePrefix = env("DB_PREFIX", "")

	return Database
}
