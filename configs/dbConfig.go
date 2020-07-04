package configs
//
// import (
//   "fmt"
//   "os"
//   "log"
//   "gopkg.in/yaml.v2"
//   "github.com/jinzhu/gorm"
//     _ "github.com/jinzhu/gorm/dialects/postgres"
// )
//
// type DatabaseConfig struct {
//   Database struct {
//     Host      string  `yaml:"host"`
//     Port      string  `yaml:"port"`
//     Adapter   string  `yaml:"adapter"`
//     User      string  `yaml:"user"`
//     Name      string  `yaml:"name"`
//     Password  string  `yaml:"password"`
//     Encoding  string  `yaml:"encoding"`
//     Pool      string  `yaml:"pool"`
//     Timeout   string  `yaml:"timeout"`
//     Sslmode   string  `yaml:"sslmode"`
//   } `yaml:"database"`
// }
//
// // ValidateConfigPath just makes sure, that the path provided is a file,
// // that can be read
// // func ValidateConfigPath(path string) error {
// //     s, err := os.Stat(path)
// //     if err != nil {
// //         return err
// //     }
// //     if s.IsDir() {
// //         return fmt.Errorf("'%s' is a directory, not a normal file", path)
// //     }
// //     return nil
// // }
//
//
// func LoadDatabaseConfig() (*DatabaseConfig, error) {
//   var configPath string = "configs/database.yml"
//   // var configPath string = "database.yml"
//   config := &DatabaseConfig{}
//   if err := ValidateConfigPath(configPath); err != nil {
//     return nil, err
//   }
//
//   file, err := os.Open(configPath)
//   if err != nil {
//     return nil, err
//   }
//   defer file.Close()
//
//   // Init new YAML decode
//   d := yaml.NewDecoder(file)
//
//   // Start YAML decoding from file
//   if err := d.Decode(&config); err != nil {
//     return nil, err
//   }
//
//   return config, nil
// }
//
//
// func LoadDatabase() (*gorm.DB, error) {
//   config, err := LoadDatabaseConfig()
//
//   if(err != nil) {
//     log.Fatalf("Unable able to access database due to %v", err)
//     return nil, err
//   }
//
//   log.Printf("--------------\n%v--------%v--------", config, config.Database.Host)
//
//   uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%q sslmode=%s",
//     config.Database.Host, config.Database.Port, config.Database.User, config.Database.Name, config.Database.Password, config.Database.Sslmode)
//
//   log.Printf("=========URI: %s", uri)
//   conn, err := gorm.Open("postgres", uri)
//   // conn, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=medi_chat_db password= sslmode=disable")
//   // conn, err := gorm.Open("postgres", "postgres://scalixte@localhost/medi_chat_db?sslmode=disable")
//
//   if(err != nil) {
//     log.Fatalf("Unable able to access database due to %v", err)
//     return nil, err
//   }
//
//   // defer conn.Close()
//   return conn, nil
// }
