package seeders

import (
	"log"

	"github.com/rifqi142/indico-be/internal/models"
	"gorm.io/gorm"
)

func RunAllSeeders(db *gorm.DB) {
	// Check if data already exists
	var count int64
	db.Model(&models.Voucher{}).Count(&count)
	
	if count > 0 {
		log.Println("Database already has data, skipping seeders...")
		return
	}
	
	log.Println("Running all seeders...")
	
	SeedVouchers(db)
	
	log.Println("All seeders completed successfully")
}
