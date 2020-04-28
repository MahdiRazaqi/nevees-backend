package main

import (
	"github.com/neveesco/nevees-backend/config"
	"github.com/neveesco/nevees-backend/database"
	"github.com/neveesco/nevees-backend/web"
)

func main() {
	config.Load()
	database.Connect()
	web.Start()
}
