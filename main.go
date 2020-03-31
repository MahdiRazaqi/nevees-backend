package main

import (
	"github.com/MahdiRazaqi/nevees-backend/database"
	"github.com/MahdiRazaqi/nevees-backend/web"
)

func main() {
	database.ConnectMongo()
	web.Start()
}
