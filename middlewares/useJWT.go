package middlewares

import (
  "fmt"
  "strings"
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/DuCalixte/MediChat-Users/helpers"
)

func UseJWT() gin.HandlerFunc {
  return func(c *gin.Context) {
    var isValid bool
    var errMsg string
    // var extractedToken string
    // token := c.Query("authToken")
    token := c.GetHeader("Authorization")

    isValid = true
    errMsg = "Unauthorized!"
    // Token should not be empty
    if token == "" {
      isValid = false
    } else {
      extractedToken := strings.Trim(strings.Split(token, "Bearer ")[1], "\"")
      if extractedToken == "" {
        isValid = false
      } else {
        _, err := helpers.ParseToken(extractedToken)
        if err != nil {
          isValid = false
          errMsg = fmt.Sprintf("Unauthorized: %v", err)
        }
      }
    }

    if !isValid {
      c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": errMsg})
      c.Abort()
      return
    }

    c.Next()
  }
}
