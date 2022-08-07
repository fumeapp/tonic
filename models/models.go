package models

import (
	"fmt"
	"log"
	"time"

	"github.com/fumeapp/skele/pkg/setting"
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
	Db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				setting.Database.User,
				setting.Database.Password,
				setting.Database.Host,
				setting.Database.Database,
			)},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
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
	Db.Create(&user)
}
