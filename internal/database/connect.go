package database

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Url struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	URL       string    `json:"url"`
	Code      string    `json:"code"`
	Clicks    int       `json:"clicks"`
	LastClick time.Time `json:"last_click"`
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	if err := db.AutoMigrate(&Url{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
