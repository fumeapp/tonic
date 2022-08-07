package models

import (
	"fmt"
	"log"
	"time"

	"github.com/fumeapp/skele/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/octoper/go-ray"
)

var Db *gorm.DB

// gorm.Model definition
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Setup() {
	var err error
	Db, err = gorm.Open(
		setting.Database.Driver,
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.Database.User,
			setting.Database.Password,
			setting.Database.Host,
			setting.Database.Database,
		),
	)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.Database.TablePrefix + defaultTableName
	}

	Db.SingularTable(true)

}

func Truncate() {
	Db.Exec("DROP TABLE user")
	Db.Exec("DROP TABLE provider")
}

func Migrate() {
	Db.AutoMigrate(&User{}, &Provider{})
}

func Seed() {
	user := User{Name: "kevin olson", Email: "acidjazz@gmail.com", Avatar: "https://avatars.githubusercontent.com/u/967369?v=4"}
	ray.Ray(user)
	Db.Create(&user)
}
