package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserAuth struct {
	gorm.Model
	UserId      string
	LastLogin   *time.Time
	AuthProfile string `gorm:"default:'local'"`
	PassHash    string `gorm:"type:varchar(100);unique_index"`
	PassSalt    string `gorm:"type:varchar(100);unique_index"`
}

func createUserAuthTable(db *gorm.DB) {
	// Create table for model `User`
	db.CreateTable(&UserAuth{})

	// userAuth := UserAuth{UserId: "1", AuthProfile: "Jinzhu", PassHash: "sce", PassSalt: "hshs"}
	// db.NewRecord(userAuth) // => returns `true` as primary key is blank
	// db.Create(&userAuth)
}

func dropUserAuthTable(db *gorm.DB) {
	// Drop model `User`'s table
	db.DropTable(&UserAuth{})
}
