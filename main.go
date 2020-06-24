package main

import (
  // "fmt"
  "log"
  "github.com/DuCalixte/MediChat-Users/models"
    // _ "github.com/DuCalixte/MediChat-Users/models/users"
)

func main() {
  db, err := LoadDatabase()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v \n- database:%v", err, db)
    return
  }
  // db, err := LoadDatabase()
  // log.Printf("Unable to connect to database with error: \nErr %v, \nDB %v", db, err)
  // defer db.Close()
  // if(!db.HasTable(&User{})){ createUserTable(db) }
  // createUserTable(db)
  dropUserTable(db)
  // defer db.Close()
}
