package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init(database *gorm.DB) {
	DB = database

	if(DB == nil) {
		log.Printf("A database should be provided")
		return
	}

	DropUserTable()

	// if(!DB.HasTable(&User{})){ CreateUserTable() }
  // DB.AutoMigrate(&User{})
	//
	// if(!DB.HasTable(&UserPref{})){ CreateUserPrefTable() }
  // DB.AutoMigrate(&UserPref{})
	//
	// if(!DB.HasTable(&UserAuth{})){ CreateUserAuthTable() }
  // DB.AutoMigrate(&UserAuth{})
	//
	// if(!DB.HasTable(&Channel{})){ CreateChannelTable() }
  // DB.AutoMigrate(&Channel{})
	//
	// if(!DB.HasTable(&UserChannel{})){ CreateUserChannelTable() }
  // DB.AutoMigrate(&UserChannel{})
}

func CloseDB() {
	defer DB.Close()
}
