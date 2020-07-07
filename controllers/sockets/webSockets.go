package sockets

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
	"github.com/DuCalixte/MediChat-Users/utilSocket"
)

// HandleConnection godoc
// @Summary Connects a user on a channel to the websocket circuit
// @Description gets User ID and Channel ID and creates a new client connection
// @USER_ID get-string-by-int
// @CHANNEL_ID get-string-by-int
// @Accept  json
// @Param user_id path int true "User ID"
// @Param id path int true "Channel ID"
// @Header 200 {string} Token "qwerty"
// @Router /user/{userId}/channel/{id} [get]
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
    socketChannel := utilSocket.BUS.At(channel.ID)
    utilSocket.HandleIncoming(&socketChannel, c.Writer, c.Request, user.Email)
  } else {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "User is not allowed in channel!"})
    return
  }
}
