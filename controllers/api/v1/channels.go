package v1

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
)

func GetChannels(c *gin.Context) {
  channels, total, err := models.GetChannels()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v", err,)
    return
  }

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": err})
    return
	}

  // total := users(&count)
  data := make(map[string]interface{})
	data["lists"] = channels
	data["total"] = total
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func GetChannel(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()

  channel, err := models.GetChannel(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": "Record not found!"})
    return
	}

  users := models.UserListChannel(channel.ID)
  channel.Users = users

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": channel})
}

func CreateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

func UpdateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
}

func DeleteChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}
