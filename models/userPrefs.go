package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/DuCalixte/MediChat-Users/helpers"
)

type UserPref struct {
	gorm.Model
	UserId  		uint		`json:"userId" gorm:"primary_key";valid:"required"`
	Firstname 	string 	`json:"firstname" gorm:"type:varchar(100)"`
	Lastname  	string 	`json:"lastname"  gorm:"type:varchar(100)"`
	Nickname  	string 	`json:"nickname" 	valid:"required" gorm:"type:varchar(100);unique_index; not null"`
	Gravatar 		string 	`json:"gravatar"`
	Color   		string 	`json:"color" valid:"required"`
	Status 			string 	`json:"status" gorm:"type:varchar(500)"`
	IsLocal 		bool   	`json:"isLocal" gorm:"default:true"`
}

type CreateUserPrefModel struct {
  UserId 		uint		`form:"userId" 			json:"userId" binding:"required"`
	Nickname 	string 	`form:"nickname" 		json:"nickname" binding:"required"`
	Firsname 	string 	`form:"firstname" 	json:"firstname" binding:"required"`
	Lastname 	string 	`form:"lastname" 		json:"lastname" binding:"required"`
  Gravatar 	string 	`form:"gravatar" 		json:"gravatar" binding:"required"`
	Status	 	string 	`form:"status" 			json:"status" binding:"required"`
	IsLocal 	string 	`form:"isLocal" 		json:"isLocal" binding:"required"`
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
		Status:			data["status"].(string),
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
