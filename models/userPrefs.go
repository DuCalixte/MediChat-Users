package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserPref struct {
	gorm.Model
	UserId  string `gorm:"primary_key";valid:"required"`
	Avatar  string
	Color   string `valid:"required"`
	AboutMe string `gorm:"type:varchar(500);unique_index"`
	IsLocal bool   `gorm:"default:true"`
}

type CreateUserPrefModel struct {
  UserId 		string	`json:"UserId" binding:"required"`
  Avatar 		string 	`json:"Avatar" binding:"required"`
	AboutMe 	string 	`json:"AboutMe" binding:"required"`
	IsLocal 	string 	`json:"IsLocal" binding:"required"`
}

func CreateUserPrefsTable() {
	// Create table for model `User`
	DB.CreateTable(&UserPref{})
}

func DropUserPrefsTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserPref{})
}

func GetUserPref(userId int) (UserPref, error) {
	var userPref UserPref
	if err := DB.Where("UserId = ?", userId).First(&userPref).Error; err != nil {
    return UserPref{}, err
  }

	return userPref, nil
}

func CreateUserPref(userId int) (UserPref, error) {
	return UserPref{}, nil
}

func UpdateUserPref(userId int) (UserPref, error) {
	return UserPref{}, nil
}
