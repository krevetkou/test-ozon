package db

import (
	"github.com/jinzhu/gorm"
	"github.com/krevetkou/test-ozon/app/models"
	"log"
	"os"
)

func OpenDB(database string) *gorm.DB {

	databaseDriver := os.Getenv("DATABASE_DRIVER")

	db, err := gorm.Open(databaseDriver, database)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if err = AutoMigrate(db); err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Post{}, &models.Comment{}).Error
}
