package main

import (
	"os"

	"github.com/MuhammedYahiya/Ecom-api/pkg/config"
	"github.com/MuhammedYahiya/Ecom-api/pkg/controller"
	"github.com/MuhammedYahiya/Ecom-api/pkg/db"
	"github.com/MuhammedYahiya/Ecom-api/pkg/server"
)

func main() {
	config.LoadConfig()
	db.ConnectDb()

	engine := server.ServerConnect()
	controller.InitializeRouter(engine)
	engine.Run(":" + os.Getenv("PORT"))
}
