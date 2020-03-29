package main

import (
	"github.com/MahdiRazaqi/nevees-backend/connection"
	"github.com/MahdiRazaqi/nevees-backend/web"
)

func main() {
	connection.ConnectMongo()
	web.Start()
}
