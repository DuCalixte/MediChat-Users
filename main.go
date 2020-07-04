// @title MediChat API
// @version 1.0
// @description A chat service api with websocket.
// @termsOfService http://swagger.io/terms/

// @contact.name MediChat API Support
// @contact.email ducalixte.@gmail.com

// @license.name MIT
// @license.url https://github.com/DuCalixte/MediChat-Users/blob/master/LICENSE
package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
  // "github.com/swaggo/gin-swagger"
  _ "github.com/swaggo/gin-swagger/example/basic/docs"
  // "github.com/DuCalixte/MediChat-Users/helpers"
  "github.com/DuCalixte/MediChat-Users/models"
    // _ "github.com/DuCalixte/MediChat-Users/models/users"
  "github.com/DuCalixte/MediChat-Users/configs"
  "github.com/DuCalixte/MediChat-Users/router"
  // "github.com/DuCalixte/MediChat-Users/hubsockets"
  // "github.com/DuCalixte/MediChat-Users/configs/routes"

  // "github.com/DuCalixte/MediChat-Users/configs"
  // _ "github.com/DuCalixte/MediChat-Users/configs/config"
  // "github.com/DuCalixte/MediChat-Users/configs"
)

func InitApp() {
  configs.InitSettings()
  InitDatabase()
  InitWebSocket()
  // helpers.InitSocket()
  // router.InitWebServer()
  //
  InitWebServer()
}

func InitWebSocket() {
  // hub := hubsockets.newHub()
	// go hub.run()
  // Allow collection of memory referenced by the caller by doing all work in
  // new goroutines.
}

func InitWebServer() {
  routesInit := router.InitRoutes()
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
}

func InitDatabase() {
  models.Init()
  models.InitializeTable()
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
