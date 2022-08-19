package main

import (
	"github.com/fumeapp/tonic/cmd"
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
)

func Init() {
	setting.Setup()
	database.Setup()
}

func main() {
	cmd.Execute()
}
