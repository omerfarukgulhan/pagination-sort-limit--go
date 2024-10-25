package postgresql

import (
	"gorm.io/gorm"
	"log"
	"pagination/domain/entities"
)

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&entities.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	} else {
		log.Println("Tables migrated successfully.")
	}
}
