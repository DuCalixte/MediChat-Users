package main

import (
	// "log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Channel struct {
	gorm.Model
	OwnerId      string
	name         string
	description  string `gorm:"type:varchar(100);unique_index"`
	IsPrivate    bool   `gorm:"default=true"`
}
