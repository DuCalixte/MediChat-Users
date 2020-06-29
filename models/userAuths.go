package models

import (
	"time"
	// "database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserAuth struct {
	gorm.Model
	UserId     	uint 				`gorm:"primary_key:UserId;unique_index"`
	LastLogin   time.Time		`sql:"DEFAULT:current_timestamp"`
	AuthProfile string 			`gorm:"default:'local'"`
	PassHash    string 			`gorm:"type:varchar(512);unique_index"`
	PassSalt    string 			`gorm:"type:varchar(512);unique_index"`
}

type AuthenticateUserModel struct {
	Password 		string 	`form:"password" 	json:"password" binding:"required"`
	Email 			string 	`form:"email"			json:"email" binding:"required"`
}

func CreateUserAuthTable() {
	// Create table for model `User`
	DB.CreateTable(&UserAuth{})
}

func DropUserAuthTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserAuth{})
}

func CreateUserAuth(user User, hash string, salt string ) (UserAuth, error) {
	 userAuth := UserAuth {
		UserId: user.ID,
		PassHash: hash,
		PassSalt: salt }

		DB.NewRecord(userAuth) // => returns `true` as primary key is blank

		if err := DB.Create(&userAuth).Error; err != nil {
			return UserAuth{}, err
		}

		return userAuth, nil
}

func GetUserAuthByUserId(userId uint) (UserAuth, error) {
	var userAuth UserAuth
	if err := DB.Where("user_id = ?", userId).First(&userAuth).Error; err != nil {
    return UserAuth{}, err
  }

	return userAuth, nil
}

func UpdateUserAuth(userId int, lastLogin *time.Time) (UserAuth, error) {
	return UserAuth{}, nil
}
