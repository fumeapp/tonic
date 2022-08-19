package tonic

import (
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
)

func Init() {
	setting.Setup()
	database.Setup()
}
