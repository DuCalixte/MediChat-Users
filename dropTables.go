package main

import (
  // "fmt"
  "log"
  // "time"
  // "net/http"
  // "github.com/DuCalixte/MediChat-Users/helpers"
  "github.com/DuCalixte/MediChat-Users/models"
    // _ "github.com/DuCalixte/MediChat-Users/models/users"
  "github.com/DuCalixte/MediChat-Users/configs"
)

func main() {
  db, err := configs.LoadDatabase()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v \n- database:%v", err, db)
    return
  }
  // models.Init(db)

  if(db.HasTable(&models.User{})){ db.DropTable(&models.User{}) }
  if(db.HasTable(&models.UserPref{})){ db.DropTable(&models.UserPref{}) }
  if(db.HasTable(&models.UserAuth{})){ db.DropTable(&models.UserAuth{}) }
  if(db.HasTable(&models.Channel{})){ db.DropTable(&models.Channel{}) }
  if(db.HasTable(&models.UserChannel{})){ db.DropTable(&models.UserChannel{}) }

  defer db.Close()
}
