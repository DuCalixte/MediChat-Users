package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}
