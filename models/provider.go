package models

type Provider struct {
	ID      uint `gorm:"primary_key" json:"id"`
	Avatar  string
	Name    string
	Payload string `gorm:"type:json"`
	UserId uint
}
