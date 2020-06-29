package api

import(
  // "log"
  "net/http"
  // "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/DuCalixte/MediChat-Users/models"
  "github.com/DuCalixte/MediChat-Users/helpers"
)

func Authorize(c *gin.Context) {
  params := models.CreateUserModel{}
  c.BindJSON(&params)

  password := params.Password
  salt, hash, err := helpers.GeneratePasswordEncrypt(password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Unable to register!"})
    return
	}

  user, err := models.CreateUser(params)
  if err != nil {
    c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "error": "Unable to create new user!"})
    return
	}

  models.CreateUserAuth(user, hash, salt)

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
