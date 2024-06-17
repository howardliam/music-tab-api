package database

import (
	"log"
	"strconv"

	"github.com/howardliam/music-tab-api/config"
	"github.com/howardliam/music-tab-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(postgresConf config.PostgresConfig) *gorm.DB {
	user := postgresConf.Username + ":" + postgresConf.Password
	address := postgresConf.Address + ":" + strconv.Itoa(int(postgresConf.Port))

	connString := "postgres://" + user + "@" + address + "/" + postgresConf.DatabaseName

	db, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Tab{})
	db.AutoMigrate(&models.Song{})
	db.AutoMigrate(&models.Album{})
	db.AutoMigrate(&models.Band{})
}
