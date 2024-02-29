package database

import (
	"log"
	"os"

	"github.com/rest/fiber-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"),&gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database \n",err.Error())
		os.Exit(2)

	}

	log.Println("Connected to the database succcefuly")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	/// add migrations
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Database = DbInstance{Db:db}
}
