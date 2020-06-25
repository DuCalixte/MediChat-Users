package configs

import (
  "log"
  // "net/http"
  "github.com/gin-gonic/gin"

  "github.com/DuCalixte/MediChat-Users/controllers/api/v1"
  // "github.com/DuCalixte/MediChat-Users/controllers/api"
  //   _ "github.com/DuCalixte/MediChat-Users/controllers/api/v1"
)

func InitRoutes() *gin.Engine {
  route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

  apiv1 := route.Group("/api/v1")

  apiv1.Use(gin.Logger())
  {
    // Users API
    apiv1.GET("/users/:id", v1.GetUser)
    apiv1.GET("/users", v1.GetUsers)

    // User Preferences API
    apiv1.GET("/userPrefs/:id", v1.GetUserPref)
    apiv1.POST("/userPrefs/:user_id", v1.CreateUserPref)
    apiv1.PUT("/userPrefs/:id", v1.UpdateUserPref)

    // Channels API
    apiv1.GET("/channels/:id", v1.GetChannel)
    apiv1.POST("/channels/:user_id", v1.CreateChannel)
    apiv1.PUT("/channels/:id", v1.UpdateChannel)
    apiv1.DELETE("/channels/:id", v1.DeleteChannel)

    apiv1.GET("/channelList/:userId", v1.GetChannelList)

  }

  return route
}

func SayHello() {
  log.Println("Hello")
}
