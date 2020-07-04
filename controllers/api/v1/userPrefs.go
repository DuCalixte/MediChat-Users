package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"

  "github.com/DuCalixte/MediChat-Users/models"
)

// GetUserPref godoc
// @Summary Show information on specific user profile
// @Description provides user profile
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "UserPref ID"
// @Success 200 {object} models.UserPref
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/userPrefs/{id} [get]
func GetUserPref(c *gin.Context) {
  // TODO Implement me!
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": models.UserPref{}})
}

// UpdateUserPref godoc
// @Summary Update an userPref
// @Description Update by json userPref
// @Tags userPrefs
// @Accept  json
// @Produce  json
// @Param  id path int true "UserPref ID"
// parameters:
//  - in: body
//    required: true
//    schema:
//       $ref: "#/definitions/UpdateUserPrefRequest"
// @Success 200 {object} models.UserPref
// @Router /api/v1/userPrefs/{id} [put]
func UpdateUserPref(c *gin.Context) {
  // TODO Implement me!
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": models.UserPref{}})
}
