package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserAuth struct {
	gorm.Model
	UserId      string 			`gorm:"primary_key";valid:"required"`
	LastLogin   *time.Time
	AuthProfile string 			`gorm:"default:'local'"`
	PassHash    string 			`gorm:"type:varchar(100);unique_index"`
	PassSalt    string 			`gorm:"type:varchar(100);unique_index"`
}

func CreateUserAuthTable(db *gorm.DB) {
	// Create table for model `User`
	DB.CreateTable(&UserAuth{})
}

func DropUserAuthTable() {
	// Drop model `User`'s table
	DB.DropTable(&UserAuth{})
}

func createUserAuth(userId int ) (UserAuth, error) {
	return UserAuth{}, nil
}

func UpdateUserAuth(userId int, lastLogin *time.Time) (UserAuth, error) {
	return UserAuth{}, nil
}
