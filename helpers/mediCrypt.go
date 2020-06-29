package helpers

import (
  "fmt"
  "log"
  "errors"
  "strings"
  "golang.org/x/crypto/bcrypt"
)

func GeneratePasswordEncrypt(password string) (string, string, error) {
  pwd := []byte(password)

  if (len(pwd) < 8) {
    log.Printf("lenth %d", len(pwd))
    fmt.Printf("password must be greater than 8 characters")
    return "", "", errors.New("Insufficient characters.")
  }

  costFactor := 11
  hash, err := bcrypt.GenerateFromPassword(pwd, costFactor)
  if err != nil {
    fmt.Printf("error generating bcrypt hash: %v\n", err)
    return "","", err
  }
  cipher := strings.SplitN(string(hash), "$", 4)[3]

  return cipher[0:22], cipher[22:], nil
}

func DecryptEncryptedGeneratedPassword(password string, hash string, salt string) (bool, error) {
  hashsalt := []byte("$2a$11$" + salt + hash)
  pwd := []byte(password)

  err := bcrypt.CompareHashAndPassword(hashsalt, pwd)
    if err != nil {
        return false, err
    }

    return true, nil
}
