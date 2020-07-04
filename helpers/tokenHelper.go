package helpers

import (
  // "fmt"
	"time"
  "github.com/dgrijalva/jwt-go"

    "github.com/DuCalixte/MediChat-Users/settings"
)

// var jwtSecret []byte

type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(email string, password string) (string, error) {
  jwtSecret   := []byte(settings.SecuritySetting.JwtSecret)
  nowTime     := time.Now()
	expireTime  := nowTime.Add(3 * time.Hour)

  claims := Claims {
    EncodeMD5(email),
    EncodeMD5(password),
    jwt.StandardClaims {
      ExpiresAt: expireTime.Unix(),
      Issuer: settings.AppSetting.Name, }}

  tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  token, err := tokenClaims.SignedString(jwtSecret)

  return token, err

}

func ParseToken(token string) (*Claims, error) {
  jwtSecret   := []byte(settings.SecuritySetting.JwtSecret)
  tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

  if tokenClaims != nil {
    if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
      return claims, nil
    }
  }

  return nil, err
}

func ValidateToken(token string) (bool) {
  var isValid bool

  isValid = true


  return isValid
}
