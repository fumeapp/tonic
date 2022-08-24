package tonic

import (
	"time"

	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
)

var Before int64

func Init() {
	Before = time.Now().UnixMicro()
	setting.Setup()
	database.Setup()
}
