package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

// func GetChannels(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
// }

func CreateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

func UpdateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
}

func DeleteChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}
