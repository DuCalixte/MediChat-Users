package models

import (
	"log"
	"fmt"
	// "reflect"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DuCalixte/MediChat-Users/settings"
)

var DB *gorm.DB

func Init() {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%q sslmode=%s",
		settings.DatabaseSetting.Host,
		settings.DatabaseSetting.Port,
		settings.DatabaseSetting.User,
		settings.DatabaseSetting.Name,
		settings.DatabaseSetting.Password,
		settings.DatabaseSetting.Sslmode)

	var err error
	DB, err = gorm.Open("postgres", uri)
	if(err != nil) {
    log.Fatalf("Unable able to access database due to %v", err)
    return
  }
}

func InitializeTable() {

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

func DropDBTables() {
	if(DB.HasTable(&User{})){ DB.DropTable(&User{}) }
	if(DB.HasTable(&UserPref{})){ DB.DropTable(&UserPref{}) }
	if(DB.HasTable(&UserAuth{})){ DB.DropTable(&UserAuth{}) }
	if(DB.HasTable(&Channel{})){ DB.DropTable(&Channel{}) }
	if(DB.HasTable(&UserChannel{})){ DB.DropTable(&UserChannel{}) }
}

func CloseDB() {
	defer DB.Close()
}
