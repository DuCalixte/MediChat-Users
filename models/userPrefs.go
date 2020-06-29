package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/DuCalixte/MediChat-Users/helpers"
)

type UserPref struct {
	gorm.Model
	UserId  	uint		`json:"userId" gorm:"primary_key";valid:"required"`
	Gravatar 	string 	`json:"gravatar"`
	Color   	string 	`json:"color" valid:"required"`
	AboutMe 	string 	`json:"aboutMe" gorm:"type:varchar(500)"`
	IsLocal 	bool   	`json:"isLocal" gorm:"default:true"`
}

type CreateUserPrefModel struct {
  UserId 		uint		`form:"userId" 		json:"userId" binding:"required"`
  Gravatar 	string 	`form:"gravatar" 	json:"gravatar" binding:"required"`
	AboutMe 	string 	`form:"aboutMe" 	json:"aboutMe" binding:"required"`
	IsLocal 	string 	`form:"isLocal" 	json:"isLocal" binding:"required"`
}

func CreateUserPrefTable() {
	// Create table for model `User`
	DB.CreateTable(&UserPref{})
}

func DropUserPrefTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserPref{})
}

func GetUserPref(userId uint) (UserPref, error) {
	var userPref UserPref
	if err := DB.Where("UserId = ?", userId).First(&userPref).Error; err != nil {
    return UserPref{}, err
  }

	return userPref, nil
}

func CreateBasicUserPref(userId uint) (UserPref, error) {
	color := helpers.RandomColor()
	fmt.Printf("color is: %s", color)
	userPref := UserPref {
		UserId:		userId,
		Color: helpers.RandomColor() }

	DB.NewRecord(userPref) // => returns `true` as primary key is blank

	if err := DB.Create(&userPref).Error; err != nil {
		return UserPref{}, err
	}

	return userPref, nil
}

func CreateUserPref(userId uint, data map[string]interface{}) (UserPref, error) {
	userPref := UserPref {
		UserId:			userId,
		Gravatar:		data["gravatar"].(string),
		Color:			data["color"].(string),
		AboutMe:		data["aboutMe"].(string),
		IsLocal:		data["isLocal"].(bool) }

	DB.NewRecord(userPref) // => returns `true` as primary key is blank

	if err := DB.Create(&userPref).Error; err != nil {
		return UserPref{}, err
	}

	return userPref, nil
}

func UpdateUserPref(userId uint) (UserPref, error) {
	return UserPref{}, nil
}
