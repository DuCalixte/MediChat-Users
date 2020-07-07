package sockets

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
	"github.com/DuCalixte/MediChat-Users/utilSocket"
)


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

    // utilSocket.HandleIncoming(utilSocket.BUS.channels[channelId], c.Writer, c.Request, user.Email)
    socketChannel := utilSocket.BUS.At(channel.ID)
    utilSocket.HandleIncoming(&socketChannel, c.Writer, c.Request, user.Email)
  } else {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "User is not allowed in channel!"})
    return
  }
}
