package models

import (
  // "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserChannel struct {
	gorm.Model
	UserId      string `gorm:"primary_key"`
	ChannelId   string `gorm:"primary_key"`
	IsActive     bool `gorm:"default:true"`
}

func CreateUserChannelsTable() {
	// Create table for model `User`
	DB.CreateTable(&UserPref{})
}

func DropUserChannelsTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserPref{})
}

func ChannelList(userId int)([]Channel, error) {
	return []Channel{}, nil
}
