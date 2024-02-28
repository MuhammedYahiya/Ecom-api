package main

import (
	"os"

	"github.com/MuhammedYahiya/Ecom-api/pkg/config"
	"github.com/MuhammedYahiya/Ecom-api/pkg/db"
	"github.com/MuhammedYahiya/Ecom-api/pkg/server"
)

func main() {
	config.LoadConfig()
	db.ConnectDb()
	engine := server.ServerConnect()
	engine.Run(":" + os.Getenv("PORT"))
}
