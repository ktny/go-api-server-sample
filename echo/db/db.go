package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ktny/go-api-server-sample/echo/config"
)

var db *gorm.DB
var err error

func Init() {
	c := config.Config.Database

	user := c.User
	password := c.Password
	dbname := c.Name

	db, err = gorm.Open("mysql", user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

func GetConnection() *gorm.DB {
	return db
}
