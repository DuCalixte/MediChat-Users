package main

import (
  // "fmt"
  "log"

  "github.com/DuCalixte/MediChat-Users/models"
  "github.com/DuCalixte/MediChat-Users/settings"
)

func main() {
  log.Printf("Intializing all settings")
  settings.InitSettings()
  log.Printf("--------------------------------")
  log.Printf("Initializing database")
  models.Init()
  log.Printf("Droping database table")
  models.DropDBTables()
  log.Printf("Now closing the database")
  models.CloseDB()
}
