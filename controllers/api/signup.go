package api

import(
  // "log"
  "net/http"
  // "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/DuCalixte/MediChat-Users/models"
  "github.com/DuCalixte/MediChat-Users/helpers"
)

// Authorize godoc
// @Summary Authorize a user to access
// @Description authorizes the user access and provides the user resource with a token
// @Tags signin
// @Accept  json
// @Produce  json
// parameters:
//  - in: body
//    required: true
//    schema:
//       $ref: "#/definitions/createAuthorizeRequest"
// @Success 200 {object} string
// @Header 200 {object} Token "qwerty"
// @Router /api/signup [post]
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

  authToken, err := helpers.GenerateToken(user.Email, password)
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Unable generate an auth token!"})
    return
  }

  // TODO - DRY VIOLATION
  channels := models.ChannelListUser(user.ID)
  user.Channels = channels
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user, "authToken": authToken})
}
