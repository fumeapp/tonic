package setting

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

type DatabaseSetting struct {
	Connection   string
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
	TablePrefix string
	Logging  string
}

type YamlDatabase struct {
	Name string `yaml:"name"`
	Driver string `yaml:"driver"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Database   string `yaml:"database"`
	User   string `yaml:"user"`
	Password   string `yaml:"password"`
	TablePrefix   string `yaml:"prefix"`
}

type YamlDatabases struct {
	DefaultConnection string         `yaml:"default"`
	Logging string `yaml:"logging"`
	Databases     []YamlDatabase `yaml:"databases"`
}

var Database = &DatabaseSetting{}

func env(key string, config string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return config
}

func Setup() *DatabaseSetting {

	godotenv.Load(".env")

	config := loadConfig[YamlDatabases]("config/database.yaml")
	fmt.Println(config.DefaultConnection)

	Database.Connection = env("DB_CONNECTION", config.DefaultConnection)
	Database.Logging = env("DB_LOGGING", config.Logging)
	dbConfig := config.Databases[slices.IndexFunc(config.Databases, func(d YamlDatabase) bool { return d.Name == Database.Connection })]
	Database.Driver = env("DB_DRIVER", dbConfig.Driver)
	Database.Host = env("DB_HOST", dbConfig.Host)
	Database.Port = env("DB_PORT", dbConfig.Port)
	Database.Database = env("DB_DATABASE", dbConfig.Database)
	Database.User = env("DB_USERNAME", dbConfig.User)
	Database.Password = env("DB_PASSWORD", dbConfig.Password)
	Database.TablePrefix = env("DB_PREFIX", dbConfig.TablePrefix)

	return Database
}

func loadConfig[T any](filename string) (config *T) {
	configFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("loadConfig() readFile err #%v", err)
	}

	err2 := yaml.Unmarshal(configFile, &config)

	if err2 != nil {
		log.Fatalf("loadConfig() Unmarshall err #%v", err)
	}

	return config
}
