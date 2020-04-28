package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
	"github.com/neveesco/nevees-backend/config"
)

// MySQL connection
var MySQL *gorm.DB

// Connect to database
func Connect() {
	db, err := gorm.Open("mysql", config.CFG.MySQL.User+":"+config.CFG.MySQL.Password+"@tcp("+config.CFG.MySQL.Host+")/"+config.CFG.MySQL.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	MySQL = db
}
