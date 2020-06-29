package configs

import (
  // "fmt"
  // "time"
  // "strings"
  // "math/rand"
  // "golang.org/x/crypto/bcrypt"
)

type EncryptConfig struct {
  Encrypt struct {
    MinCount      string  `yaml:"minCount"`
    CostFactor    string  `yaml:"port"`
  } `yaml:"encrypt"`
}

func LoadEncryptConfigs() (*EncryptConfig, error) {
  // file, err := LoadConfigFile("configs/database.yml")
  configPath := "configs/encrypt.yml"
  config := &EncryptConfig{}
  decoder, err := LoadConfigFile(configPath)
  if err != nil {
    return nil, err
  }

  // Start YAML decoding from file
  if err := decoder.Decode(&config); err != nil {
    return nil, err
  }

  return config, nil
}


// func GeneratePasswordEncrypt(password string) (string, string, error) {
//   pwd := []byte(password)
//   // configs, err = LoadEncryptConfigs()
//   if config, err := LoadEncryptConfigs(); err != nil {
//     return "",""
//   }
//
//   minCount := config.Encrypt.MinCount
//   costFactor := config.Encrypt.CostFactor
//
//   hash, err := bcrypt.GenerateFromPassword(pwd, costFactor)
//   if err != nil {
//     fmt.Printf("error generating bcrypt hash: %v\n", err)
//     return "","", err
//   }
//
//   return "","", nil
// }
