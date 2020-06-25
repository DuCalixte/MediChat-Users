package models

import (
	// "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Channel struct {
	gorm.Model
	OwnerId      int		`gorm:"default":0`
	name         string	`gorm:"type:varchar(100)"`
	description  string `gorm:"type:varchar(512)"`
	IsPrivate    bool   `gorm:"default"=true`
	IsAChatBot	 bool		`gorm:"default"=false`
}


func CreateChannelTable() {
	// Create table for model `Channel`
	DB.CreateTable(&Channel{})
}

func DropChannelTable() {
	// Drop model `User`'s table
	DB.DropTable(&Channel{})
}

func GetChannel(id int)(Channel, error) {
	return Channel{}, nil
}

func CreateChannel(userId int, maps interface{})(Channel, error) {
	return Channel{}, nil
}

func UpdateChannel(id int, maps interface{})(Channel, error) {
	return Channel{}, nil
}

func CreateGalleryChat() (Channel, error) {
	return Channel{}, nil
}

func CreateAChatBot(userId int)(Channel, error) {
	return Channel{}, nil
}
