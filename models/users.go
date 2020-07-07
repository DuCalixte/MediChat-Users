package models

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/DuCalixte/MediChat-Users/helpers"
)

const (
	personnelChatDesc = "You have your own personnel channet - So go talk to your self"
)

type User struct {
	gorm.Model
	Email     	string 			`json:"email" valid:"email;required" gorm:"type:varchar(100);unique_index" email:"required"`
	UserPref		UserPref		`json:"profile" gorm:"PRELOAD:true;ForeignKey:UserId;AssociationForeignKey:ID"`
	Channels		[]*Channel	`json:"channels" gorm:"PRELOAD:true;many2many:user_channels;"`
}

type CreateUserModel struct {
  Firstname 	string 	`form:"firstname" json:"firstname" binding:"required"`
	Lastname 		string 	`form:"lastname"	json:"lastname" binding:"required"`
	Nickname 		string 	`form:"nickname"  json:"nickname" binding:"required"`
	Password 		string 	`form:"password" 	json:"password" binding:"required"`
	Email 			string 	`form:"email"			json:"email" binding:"required"`
}

func CreateUserTable() {
	// Create table for model `User`
	DB.CreateTable(&User{})
}

func CreateFirstUsers() {
	user1 := User {
		Email: "jiglesias@example.com",
		UserPref: UserPref {
			Firstname: "Julio",
			Lastname: "Iglesia",
			Nickname: "jiglesias",
			Gravatar: helpers.RandomLegoGravatar(),
			Color: helpers.RandomColor() }}
	DB.NewRecord(user1) // => returns `true` as primary key is blank
	DB.Create(&user1)

	user2 := User{
		Email: "blee@example.com",
		UserPref: UserPref {
			Firstname: "Bruce",
			Lastname: "Lee",
			Nickname: "blee",
			Gravatar: helpers.RandomLegoGravatar(),
			Color: helpers.RandomColor() }}
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
	}

}

func CreateUser(data CreateUserModel) (User, error) {
	user := User {
		Email:			data.Email,
		UserPref: UserPref {
			Firstname:	data.Firstname,
			Lastname:		data.Lastname,
			Nickname:		data.Nickname,
			Gravatar:		helpers.RandomLegoGravatar(),
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

	selfChannel, err := CreateChannel(user, data.Nickname, personnelChatDesc, true)
	if err == nil {
		DB.Model(&user).Association("Channels").Append(selfChannel)
	}

	return user, nil
}

func GetUsers() ([]*User, int, error) {
	var users []*User
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
	if err := DB.Preload("UserPref").Where("id = ?", id).First(&user).Error; err != nil {
    return User{}, err
  }

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	if err := DB.Preload("UserPref").Where("Email = ?", email).First(&user).Error; err != nil {
    return User{}, err
  }

	return user, nil
}

func GetUserWithPref(user User, userPref UserPref) User {
	DB.Model(&user).Related(&userPref, "UserPref").First(&user)
	return user
}

func DropUserTable() {
	// Drop model `User`'s table
	DB.DropTable(&User{})

	// drop related tables
	DropUserAuthTable()
	DropUserPrefTable()
}
