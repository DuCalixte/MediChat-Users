package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetChannelList(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

// func GetUserChannels(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
// }
//
// func CreateUserChannel(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
// }
//
// func UpdateUserChannel(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
// }
//
// func DeleteUserChannel(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
// }
