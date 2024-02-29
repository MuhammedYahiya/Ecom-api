package db

import (
	"fmt"
	"log"
	"os"

	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	host := os.Getenv("DBHOST")
	dbUserName := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUserName, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	log.Println("Connected to database successfully")

	DB.AutoMigrate(&domain.User{})

}
