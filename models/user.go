package models

type User struct {
	Email     string `gorm:"unique"`
	Name      string
	Avatar    string
	Providers []Provider
	Model
}
