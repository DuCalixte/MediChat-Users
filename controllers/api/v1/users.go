package v1

import (
  "net/http"
  "log"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
)

// GetUser godoc
// @Summary Show information on specific user
// @Description provides user information with preferences in models.UserPref and channels from models.Channel
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()

  user, err := models.GetUser(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": "Record not found!"})
    return
	}

  // TODO - DRY VIOLATION
  channels := models.ChannelListUser(user.ID)
  user.Channels = channels
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// GetUsers godoc
// @Summary List all Users
// @Description get all users
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} models.User
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
  users, total, err := models.GetUsers()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v", err,)
    return
  }

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": err})
    return
	}

  data := make(map[string]interface{})
	data["lists"] = users
	data["total"] = total
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
