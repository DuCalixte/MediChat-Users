package models

import (
	// "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Firstname string `json:"Firstname";valid:"required"`
	Lastname  string `json:"Lastname";valid:"required"`
	Username  string `json:"Username";gorm:"type:varchar(100);valid:"required";unique_index"`
	Email     string `json:"Email";gorm:"type:varchar(100);email:"required";unique_index"`
}

type CreateUserModel struct {
  Firstname 	string `json:"Firstname" binding:"required"`
  Lastname 		string 	`json:"Lastname" binding:"required"`
	Username 		string 	`json:"Username" binding:"required"`
	Email 			string 	`json:"Lastname" binding:"required" email:"required"`
	Password 		string 	`json:"Password" binding:"required" password:"required"`
}

func CreateUserTable() {
	// Create table for model `User`
	DB.CreateTable(&User{})
}

func DropUserTable() {
	// Drop model `User`'s table
	DB.DropTable(&User{})

	// drop related tables
	DropUserAuthTable()
	DropUserPrefsTable()
}

func CreateFirstUsers() {
	user1 := User{Firstname: "Jinzhu", Lastname: "Boo", Username: "sboo", Email: "sboo@example.com"}
	DB.NewRecord(user1) // => returns `true` as primary key is blank
	DB.Create(&user1)

	user2 := User{Firstname: "Jinzhu", Lastname: "Moo", Username: "smoo", Email: "smoo@example.com"}
	DB.NewRecord(user2) // => returns `true` as primary key is blank
	DB.Create(&user2)
}

func GetUsers() ([]*User, int, error) {
	var users []*User
	err := DB.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	var total int
	DB.Model(&users).Count(&total)

	return users, total, nil
}

func GetUser(id int) (User, error) {
	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
    return User{}, err
  }

	return user, nil
}

// func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
// 	var users []*User
// 	err := DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
//
// 	return users, nil
// }
