package helpers

import (
  "fmt"
  "log"
  "errors"
  "strings"
  "golang.org/x/crypto/bcrypt"

  "github.com/DuCalixte/MediChat-Users/settings"
)

func GeneratePasswordEncrypt(password string) (string, string, error) {
  pwd := []byte(password)

  if (len(pwd) < settings.SecuritySetting.MinPasswordSize) {
    log.Printf("length %d", len(pwd))
    fmt.Printf("password must be greater than %d characters", settings.SecuritySetting.MinPasswordSize)
    return "", "", errors.New("Insufficient characters.")
  }

  hash, err := bcrypt.GenerateFromPassword(pwd, settings.SecuritySetting.CostFactor)
  if err != nil {
    fmt.Printf("error generating bcrypt hash: %v\n", err)
    return "","", err
  }
  cipher := strings.SplitN(string(hash), "$", 4)[3]

  return cipher[0:22], cipher[22:], nil
}

func DecryptEncryptedGeneratedPassword(password string, hash string, salt string) (bool, error) {
  cypherHash := fmt.Sprintf("$2a$%d$%s%s", settings.SecuritySetting.CostFactor, salt, hash)
  hashsalt := []byte(cypherHash)
  pwd := []byte(password)

  err := bcrypt.CompareHashAndPassword(hashsalt, pwd)
    if err != nil {
        return false, err
    }

    return true, nil
}
