package router

import (
  "log"
  "fmt"
  "net/http"

  "github.com/DuCalixte/MediChat-Users/settings"
)

func InitWebServer() {
  routesInit := InitRoutes()
  readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
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
