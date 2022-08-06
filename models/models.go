package models

import (
	"fmt"
	"log"
	"time"

	"github.com/fumeapp/skele/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedOn time.Time `json:"deleted_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open(
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

	db.SingularTable(true)
}