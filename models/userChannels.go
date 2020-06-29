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
	IsActive    bool `gorm:"default:true"`
}

func CreateUserChannelTable() {
	// Create table for model `User`
	DB.CreateTable(&UserChannel{})
}

func DropUserChannelTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserChannel{})
}

func ChannelListUser(userId uint) []*Channel {
	channels := []*Channel{}
	query := "Select * from channels inner join user_channels on channels.id = user_channels.channel_id where user_channels.user_id = ?"

	DB.Raw(query, userId).Scan(&channels)
	return channels
}

func UserListChannel(channelId uint)[]*User {
	users := []*User{}
	query := "Select * from users inner join user_channels on users.id = user_channels.user_id where user_channels.channel_id = ?"

	DB.Raw(query, channelId).Scan(&users)
	return users
}
