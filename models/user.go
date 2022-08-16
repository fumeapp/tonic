package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Email     string `gorm:"unique"`
	Name      string
	Avatar    string
	Providers []Provider
	Model
	URL string `gorm:"-:all"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.URL = fmt.Sprintf("/user/%d", u.ID)
	return
}