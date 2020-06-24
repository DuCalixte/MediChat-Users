package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserPref struct {
	gorm.Model
	UserId  string
	Avatar  string
	Color   string
	AboutMe string `gorm:"type:varchar(500);unique_index"`
	isLocal bool   `gorm:"default:true"`
}

func createUserPrefsTable(db *gorm.DB) {
	// Create table for model `User`
	db.CreateTable(&UserPref{})

	// userPref := UserPref{UserId: "1", Avatar: "Jinzhu", AboutMe: "sce", Color: randomColor() }
	// db.NewRecord(userPref) // => returns `true` as primary key is blank
	// db.Create(&userPref)
}

func dropUserPrefsTable(db *gorm.DB) {
	// Drop model `User`'s table
	db.DropTable(&UserPref{})
}
