package database

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/fumeapp/tonic/setting"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var Os *opensearch.Client

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Database.Username,
		setting.Database.Password,
		setting.Database.Host,
		setting.Database.Database,
	)
}

func DURL() string {
	return "mysql://" + DSN()
}

func Setup() {
	var err error
	var logMode = logger.Error
	if setting.Database.Logging == "true" {
		logMode = logger.Info
	}
	Db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: DSN()},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		log.Fatalf("gorm.DB err: %v", err)
	}

	Os, err = opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{setting.Opensearch.Address},
		Username:  setting.Opensearch.Username,
		Password:  setting.Opensearch.Password,
	})

	if err != nil {
		log.Fatalf("opensearch.NewClient err: %v", err)
	}
}

func Truncate() {
	Db.Exec("DROP TABLE providers")
	Db.Exec("DROP TABLE users")
}

func Migrate() {
	Db.Exec("SQL FOR MIGRATION")
}

func Seed() {
	/*
		type FakeUser struct {
			Name   string `faker:"name"`
			Email   string `faker:"email"`
			Avatar string
		}
		users := []User{}
		for i := 0; i < 10; i++ {
			fakeUser := FakeUser{}
			faker.FakeData(&fakeUser)
			var fakeAvatar = "http://i.pravatar.cc/150?u=" + fakeUser.Email
			fakeUser.Avatar = fakeAvatar

			user := User{}
			user.Name = fakeUser.Name
			user.Email = fakeUser.Email
			user.Avatar = fakeUser.Avatar
			user.Providers = []Provider{
				{
					Name: "google",
					Avatar: user.Avatar,
					Payload: "{\"id\":\"12345\",\"name\":\"John Smith\"}",
				},
			}
			// Db.Create(&user)
			users = append(users, user)
		}
		Db.Create(&users)
		user := User{
			Name: "kevin olson",
			Email: "acidjazz@gmail.com",
			Avatar: "https://avatars.githubusercontent.com/u/967369?v=4",
			Providers: []Provider{
				{
					Name: "google",
					Avatar: "https://avatars.githubusercontent.com/u/967369?v=4",
					Payload: "{\"id\":\"12345\",\"name\":\"kevin olson\"}",
				},
			},

		}
	*/
}
