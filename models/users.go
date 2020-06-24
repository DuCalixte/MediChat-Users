package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Username  string `gorm:"type:varchar(100);unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
}

func createUserTable(db *gorm.DB) {
	log.Printf("Unable to connect to database with error: %v", db.DB())
	// Create table for model `User`
	db.CreateTable(&User{})

	// create related table
	createUserAuthTable(db)
	createUserPrefsTable(db)
}

func dropUserTable(db *gorm.DB) {
	// Drop model `User`'s table
	db.DropTable(&User{})

	// // drop related tables
	dropUserAuthTable(db)
	dropUserPrefsTable(db)
}
