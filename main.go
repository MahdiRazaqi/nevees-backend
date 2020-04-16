package main

import (
	"github.com/MahdiRazaqi/nevees-backend/config"
	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/MahdiRazaqi/nevees-backend/web"
)

func main() {
	config.Load()
	database.Connect()
	web.Start()
}
