package models

import (
	// "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/DuCalixte/MediChat-Users/helpers"
)

type User struct {
	gorm.Model
	Firstname 	string 			`json:"firstname" gorm:"type:varchar(100)"`
	Lastname  	string 			`json:"lastname"  gorm:"type:varchar(100)"`
	Username  	string 			`json:"username" valid:"required" gorm:"type:varchar(100);unique_index; not null"`
	Email     	string 			`json:"email" valid:"email;required" gorm:"type:varchar(100);unique_index" email:"required"`
	UserPref		UserPref		`json:"profile" gorm:"PRELOAD:true;ForeignKey:UserId;AssociationForeignKey:ID"`
	Channels		[]*Channel	`json:"channels" gorm:"PRELOAD:true;many2many:user_channels;"`
}

type CreateUserModel struct {
  Firstname 	string 	`form:"firstname" json:"firstname" binding:"required"`
	Lastname 		string 	`form:"lastname"	json:"lastname" binding:"required"`
	Username 		string 	`form:"username"  json:"username" binding:"required"`
	Password 		string 	`form:"password" 	json:"password" binding:"required"`
	Email 			string 	`form:"email"			json:"email" binding:"required"`
}

func CreateUserTable() {
	// Create table for model `User`
	DB.CreateTable(&User{})
}

func CreateFirstUsers() {
	user1 := User{Firstname: "Julio", Lastname: "Iglesia", Username: "jiglesias", Email: "jiglesias@example.com"}
	DB.NewRecord(user1) // => returns `true` as primary key is blank
	DB.Create(&user1)

	user2 := User{Firstname: "Bruce", Lastname: "Lee", Username: "blee", Email: "blee@example.com"}
	DB.NewRecord(user2) // => returns `true` as primary key is blank
	DB.Create(&user2)

	gallery, err := GetChannelGallery()
	if err == nil {
		DB.Model(&user1).Association("Channels").Append(gallery)
		DB.Model(&user2).Association("Channels").Append(gallery)
	}

	chatbotChannel, err := CreateAChatBot(user1.ID)
	if err == nil {
		DB.Model(&user1).Association("Channels").Append(chatbotChannel)
		// DB.Model(&user2).Association("Channels").Append(chatbotChannel)
		// DB.Model(&user).Association("Channels").Append([]*Channel{&chatbotChannel})
	}

}

func CreateUser(data CreateUserModel) (User, error) {
	user := User {
		Firstname:	data.Firstname,
		Lastname:		data.Lastname,
		Username:		data.Username,
		Email:			data.Email,
		UserPref: UserPref {
			Color: helpers.RandomColor() }}

	DB.NewRecord(user) // => returns `true` as primary key is blank

	if err := DB.Create(&user).Error; err != nil {
		return User{}, err
	}

	galleryChannel, err := GetChannelGallery()
	if err == nil {
		DB.Model(&user).Association("Channels").Append(galleryChannel)
	}

	chatbotChannel, err := CreateAChatBot(user.ID)
	if err == nil {
		DB.Model(&user).Association("Channels").Append(chatbotChannel)
	}
	
	return user, nil
}

func GetUsers() ([]*User, int, error) {
	var users []*User
	// err := DB.Find(&users).Error
	err := DB.Preload("UserPref").Find(&users).Error
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

func GetUserByEmail(email string) (User, error) {
	var user User
	if err := DB.Where("Email = ?", email).First(&user).Error; err != nil {
    return User{}, err
  }

	return user, nil
}

func GetUserWithPref(user User, userPref UserPref) User {
	DB.Model(&user).Related(&userPref, "UserPref").First(&user)
	return user
}

// func CreateUser(data map[string]interface{}) (User, error) {
// 	user := User {
// 		Firstname:	data["firstname"].(string),
// 		Lastname:		data["lastname"].(string),
// 		Username:		data["username"].(string),
// 		Email:			data["email"].(string) }
//
// 	DB.NewRecord(user) // => returns `true` as primary key is blank
//
// 	if err := DB.Create(&user).Error; err != nil {
// 		return User{}, err
// 	}
//
// 	return user, nil
// }

// func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
// 	var users []*User
// 	err := DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
//
// 	return users, nil
// }

func DropUserTable() {
	// Drop model `User`'s table
	DB.DropTable(&User{})

	// drop related tables
	DropUserAuthTable()
	DropUserPrefTable()
}
