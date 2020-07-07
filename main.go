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
  _ "github.com/swaggo/gin-swagger/example/basic/docs"
  "github.com/DuCalixte/MediChat-Users/models"
  "github.com/DuCalixte/MediChat-Users/router"
  "github.com/DuCalixte/MediChat-Users/utilSocket"
  "github.com/DuCalixte/MediChat-Users/settings"
  "github.com/DuCalixte/MediChat-Users/notifier"
)

func InitApp() {
  settings.InitSettings()
  utilSocket.InitBus()
  InitDatabase()
  InitWebServer()
  notifier.InitNotifier()
}

func InitWebServer() {
  routesInit := router.InitRoutes()
	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
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
}
