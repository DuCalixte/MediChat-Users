package settings

import (
  "log"
	"time"

  "github.com/go-ini/ini"
)

type App struct {
  Name       string
}
var AppSetting = &App{}

type Notification struct {
  PrivateKey       string
  PublicKey       string
  Subscriber       string
}
var NotificationSetting = &Notification{}

type Security struct {
  CostFactor      int
  JwtSecret       string
  MinPasswordSize int
  TokenDuration   time.Duration
}
var SecuritySetting = &Security{}

type Server struct {
  Host         string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
  Host      string
  Port      string
  Adapter   string
  User      string
  Name      string
  Password  string
  Encoding  string
  Pool      string
  Timeout   time.Duration
  Sslmode   string
}
var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}
var RedisSetting = &Redis{}

var cfg *ini.File
func InitSettings() {
  var err error
	cfg, err = ini.Load("./app.ini")
	if err != nil {
		log.Fatalf("settings.Initsettings, fail to parse 'settings/app.ini': %v", err)
	}

  mapTo("app", AppSetting)
  mapTo("security", SecuritySetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("notification", NotificationSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
