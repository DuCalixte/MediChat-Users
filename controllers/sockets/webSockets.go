package sockets

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
  // "github.com/DuCalixte/MediChat-Users/helpers"
  "github.com/DuCalixte/MediChat-Users/hubsockets"
)

// func HandleConnection(c *gin.Context) {
//   id := com.StrTo(c.Param("id")).MustInt()
//
//   channel, err := models.GetChannel(id)
//   if err != nil {
//     c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Unable to recognize request!"})
//     return
// 	}
//   helpers.HandleIncoming(c.Writer, c.Request, channel.IsAChatBot)
// }

func HandleConnection(c *gin.Context) {
  channelId := com.StrTo(c.Param("id")).MustInt()
  userId := com.StrTo(c.Param("userId")).MustInt()

  channel, err := models.GetChannel(channelId)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Unable to recognize request!"})
    return
	}

  user, err := models.GetUser(userId)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Unable to recognize request!"})
    return
	}

  if models.VerifyUserChannel(userId, channelId) {
    hub := hubsockets.NewHub(channel.ID)
  	go hub.Run()

    hubsockets.HandleIncoming(hub, c.Writer, c.Request, channel, user)
    // hubsockets.HandleIncoming(hub, c.Writer, c.Request)
  } else {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "User is not allowed in channel!"})
    return
  }
}
