package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetUserPref(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

func CreateUserPref(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

func UpdateUserPref(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
}
