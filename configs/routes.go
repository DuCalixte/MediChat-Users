package configs

import (
  "log"
  // "time"
  // "net/http"
  // "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/swaggo/gin-swagger"
  "github.com/swaggo/gin-swagger/swaggerFiles"

  "github.com/DuCalixte/MediChat-Users/controllers/api"
  "github.com/DuCalixte/MediChat-Users/controllers/api/v1"

  "github.com/DuCalixte/MediChat-Users/middlewares"
)

func InitRoutes() *gin.Engine {
  route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
  route.Use(middlewares.CORSMiddleware())

  apiv1 := route.Group("/api/v1")

  route.POST("/signin", api.Authenticate)
  route.POST("/signup", api.Authorize)
  // url := ginSwagger.URL("http://localhost:8001/swagger/swagger.json") // The url pointing to API definition
	// route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// TODO Use JWT
  apiv1.Use(gin.Logger())
  {
    // Users API
    apiv1.GET("/users/:id", v1.GetUser)
    apiv1.GET("/users", v1.GetUsers)

    // User Preferences API
    apiv1.GET("/userPrefs/:id", v1.GetUserPref) // TODO - Implement
    apiv1.POST("/userPrefs/:user_id", v1.CreateUserPref) // TODO - Implement
    apiv1.PUT("/userPrefs/:id", v1.UpdateUserPref) // TODO - Implement

    // Channels API
    apiv1.GET("/channels/:id", v1.GetChannel)
    apiv1.GET("/channels", v1.GetChannels)
    apiv1.POST("/channels/:user_id", v1.CreateChannel)  // TODO - Implement
    apiv1.PUT("/channels/:id", v1.UpdateChannel)  // TODO - Implement
    apiv1.DELETE("/channels/:id", v1.DeleteChannel) // TODO - Implement

  }

  return route
}

func SayHello() {
  log.Println("Hello")
}
