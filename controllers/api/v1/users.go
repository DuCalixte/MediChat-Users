package v1

import (
  "net/http"
  "log"
  "github.com/gin-gonic/gin"
  "github.com/unknwon/com"

  "github.com/DuCalixte/MediChat-Users/models"
)

func GetUser(c *gin.Context) {
  id := com.StrTo(c.Param("id")).MustInt()
  // log.Printf(c)

  user, err := models.GetUser(id)
  if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err, "data": nil})
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": "Record not found!"})
    return
	}
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func GetUsers(c *gin.Context) {
  users, total, err := models.GetUsers()
  if err != nil {
    log.Printf("Unable to connect to database with \n-error:%v", err,)
    return
  }

  if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err, "data": nil})
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": err})
    return
	}

  // total := users(&count)
  data := make(map[string]interface{})
	data["lists"] = users
	data["total"] = total
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}