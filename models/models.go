package models

import (
	"log"
	// "reflect"
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

	if(!DB.HasTable(&User{})){ CreateUserTable() }
  DB.AutoMigrate(&User{})

	if(!DB.HasTable(&UserPref{})){ CreateUserPrefTable() }
  DB.AutoMigrate(&UserPref{})

	if(!DB.HasTable(&UserAuth{})){ CreateUserAuthTable() }
  DB.AutoMigrate(&UserAuth{})

	if(!DB.HasTable(&Channel{})){ CreateChannelTable() }
  DB.AutoMigrate(&Channel{})

	if(!DB.HasTable(&UserChannel{})){ CreateUserChannelTable() }
	 DB.AutoMigrate(&UserChannel{})
	// DB.DropTable(&UserChannel{})

	// Creating Channels
	// if(reflect.DeepEqual(GalleryChannel, Channel{})){
	count := 0
	DB.Model(&Channel{}).Count(&count)
	if(count == 0){
		if err := CreateGalleryChat(); err != nil {
			log.Printf("Unable to create a basic channel")
		}
	}
	userCount := 0
	DB.Model(&User{}).Count(&userCount)
	if (userCount == 0) { CreateFirstUsers() }

}

func CloseDB() {
	defer DB.Close()
}
