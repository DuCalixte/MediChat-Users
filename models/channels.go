package models

import (
	// "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Channel struct {
	gorm.Model
	OwnerId      uint			`json:"ownerId" gorm:"default:0"`
	Name         string		`json:"name" gorm:"type:varchar(100)"`
	Description  string 	`json:"description" gorm:"type:varchar(512)"`
	IsPrivate    bool   	`json:"isPrivate" gorm:"default:true"`
	IsAChatBot	 bool			`json:"chatBot" gorm:"default:false"`
	// WebSocket		 string		`json:"websocket"  gorm:"type:varchar(256)"`
	Users				[]*User		`json:"users" gorm:"PRELOAD:true;many2many:user_channels;"`
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
	var channel Channel
	// if err := DB.Preload("User").Preload("UserPref").Where("id = ?", id).First(&channel).Error; err != nil {
	if err := DB.Where("id = ?", id).First(&channel).Error; err != nil {
    return Channel{}, err
  }

	return channel, nil
}

func GetChannelGallery() (Channel, error) {
	var channel Channel
	if err := DB.Where("owner_id = ? AND name = ?", 0, "Gallery").First(&channel).Error; err != nil {
    return Channel{}, err
  }

	return channel, nil
}

func GetChannels() ([]*Channel, int, error) {
	var channels []*Channel
	err := DB.Find(&channels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	var total int
	DB.Model(&channels).Count(&total)

	return channels, total, nil
}

func CreateChannel(user User, name string, description string, isPrivate bool)(Channel, error) {
	channel := Channel{
		OwnerId: user.ID,
		Name: name,
		IsPrivate: isPrivate,
		Description: description}
	DB.NewRecord(channel) // => returns `true` as primary key is blank
	if err := DB.Create(&channel).Error; err != nil {
		return Channel{}, err
	}
	return channel, nil
}

func CreateGalleryChat() error {
	channel := Channel{Name: "Gallery", Description: "The main channel for all users"}
	DB.NewRecord(channel) // => returns `true` as primary key is blank
	if err := DB.Create(&channel).Error; err != nil {
		return err
	}

return nil
}

func CreateAChatBot(userId uint)(Channel, error) {
	channel := Channel{
		OwnerId: userId,
		Name: "ChatBot",
		IsAChatBot: true,
		Description: "Your personal android to chat with, ask Away!!!:-)"}
	DB.NewRecord(channel) // => returns `true` as primary key is blank
	if err := DB.Create(&channel).Error; err != nil {
		return Channel{}, err
	}
	return channel, nil
}

func UpdateChannel(id int, maps interface{})(Channel, error) {
	return Channel{}, nil
}
