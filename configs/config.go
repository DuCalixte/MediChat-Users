package configs

import (
  "fmt"
  "os"
  "gopkg.in/yaml.v2"
)

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
    s, err := os.Stat(path)
    if err != nil {
        return err
    }
    if s.IsDir() {
        return fmt.Errorf("'%s' is a directory, not a normal file", path)
    }
    return nil
}

// func LoadConfigFile(filePath string) (*os.File, error) {
func LoadConfigFile(filePath string) (*yaml.Decoder, error) {
  if err := ValidateConfigPath(filePath); err != nil {
    return nil, err
  }

  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }

  defer file.Close()
  // Init new YAML decode
  decoder := yaml.NewDecoder(file)

  return decoder, nil
}

func CloseFile(file *os.File) {
  // err = defer file.Close()
  // return err
  // defer file.Close()
}
//
// func DecodeFile(file)
