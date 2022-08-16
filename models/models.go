package models

import (
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/fumeapp/tonic/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Setup() {
	var err error
	var logMode = logger.Error
	if (setting.Database.Logging == "true") {
		logMode = logger.Info
	}
	Db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				setting.Database.Username,
				setting.Database.Password,
				setting.Database.Host,
				setting.Database.Database,
			)},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

}

func Truncate() {
	Db.Exec("DROP TABLE providers")
	Db.Exec("DROP TABLE users")
}

func Migrate() {
	Db.AutoMigrate(&User{}, &Provider{})
}

func Seed() {
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
	/*
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
