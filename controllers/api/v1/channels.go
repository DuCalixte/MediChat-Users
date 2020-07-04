package v1

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
)


// GetChannel godoc
// @Summary Show information on specific channel
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Channel ID"
// @Success 200 {object} models.Channel
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/channels/{id} [get]
func GetChannel(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()

  channel, err := models.GetChannel(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Record not found!"})
    return
	}

  users := models.UserListChannel(channel.ID)
  channel.Users = users

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": channel})
}

// GetChannels godoc
// @Summary List channels
// @Description get channels
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} models.Channel
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/channels [get]
func GetChannels(c *gin.Context) {
  channels, total, err := models.GetChannels()

  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v", err)
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err})
    return
  }

  // total := users(&count)
  data := make(map[string]interface{})
	data["lists"] = channels
	data["total"] = total
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

// CreateChannel godoc
// @Summary Create channel
// @Description create a new channel
// @Tags channels
// @Accept  json
// @Produce  json
// parameters:
//  - in: body
//    required: true
//    schema:
//       $ref: "#/definitions/createChannelRequest"
// @Success 200 {object} models.Channel
// @Header 200 {object} Token "qwerty"
// @Router /api/v1/channels [post]
func CreateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}

// UpdateChannel godoc
// @Summary Update an channel
// @Description Update by json channel
// @Tags channels
// @Accept  json
// @Produce  json
// @Param  id path int true "Channel ID"
// parameters:
//  - in: body
//    required: true
//    schema:
//       $ref: "#/definitions/UpdateChannelRequest"
// @Success 200 {object} models.Channel
// @Router /api/v1/channels/{id} [put]
func UpdateChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 5})
}

// DeleteChannel godoc
// @Summary Delete a channel
// @Description Delete by channel ID
// @Tags channels
// @Accept  json
// @Produce  json
// @Param  id path int true "Channel ID" Format(int64)
// @Success 204 {object} models.Channel
func DeleteChannel(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": 0})
}
