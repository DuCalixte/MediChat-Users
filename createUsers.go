package loco

import (
  "log"
  "github.com/DuCalixte/MediChat-Users/configs"
  "github.com/DuCalixte/MediChat-Users/models"
)

func InitDatabase() {
  db, err := configs.LoadDatabase()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v \n- database:%v", err, db)
    return
  }
  models.Init(db)
}

func main() {
  InitDatabase()
  models.CreateFirstUsers()
}
