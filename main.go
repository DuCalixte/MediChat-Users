package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
  // "github.com/DuCalixte/MediChat-Users/helpers"
  "github.com/DuCalixte/MediChat-Users/models"
    // _ "github.com/DuCalixte/MediChat-Users/models/users"
  "github.com/DuCalixte/MediChat-Users/configs"
  // "github.com/DuCalixte/MediChat-Users/configs/routes"

  // "github.com/DuCalixte/MediChat-Users/configs"
  // _ "github.com/DuCalixte/MediChat-Users/configs/config"
  // "github.com/DuCalixte/MediChat-Users/configs"
)

func InitApp() {
  InitDatabase()
  InitWebServer()
}

func InitDatabase() {
  db, err := configs.LoadDatabase()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v \n- database:%v", err, db)
    return
  }
  models.Init(db)
}

func InitWebServer() {
  routesInit := configs.InitRoutes()
	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := fmt.Sprintf(":%d", 8001)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routesInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

  // configs.SayHello()
}

func main() {
  InitApp()
  // db, err := LoadDatabase()
  // log.Printf("Unable to connect to database with error: \nErr %v, \nDB %v", db, err)
  // defer db.Close()
  // if(!db.HasTable(&User{})){ createUserTable(db) }
  // createUserTable(db)
  // dropUserTable(db)
  // defer db.Close()
}
