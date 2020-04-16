package database

import (
	"log"

	"github.com/MahdiRazaqi/nevees-backend/database/mysql"
)

// Connect to all databases
func Connect() {
	if err := mysql.Connect(); err != nil {
		log.Fatal("connecting to mysql ", err)
	}
}
