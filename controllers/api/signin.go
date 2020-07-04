package api

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/DuCalixte/MediChat-Users/models"
  "github.com/DuCalixte/MediChat-Users/helpers"
)

func Authenticate(c *gin.Context) {
  params := models.AuthenticateUserModel{}
  c.BindJSON(&params)

  password := params.Password
  email    := params.Email

  user, err := models.GetUserByEmail(email)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Unable access information!"})
    return
	}

  userAuth, err := models.GetUserAuthByUserId(user.ID)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Unable access information!"})
    return
	}

  found, err := helpers.DecryptEncryptedGeneratedPassword(password, userAuth.PassHash, userAuth.PassSalt)
  if (found == false || err != nil) {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Unable access information!"})
    return
	}

  authToken, err := helpers.GenerateToken(email, password)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Unable generate an auth token!"})
    return
	}

  // TODO - DRY VIOLATION
  channels := models.ChannelListUser(user.ID)
  user.Channels = channels
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user, "authToken": authToken})
}
