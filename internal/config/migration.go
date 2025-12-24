package config

import (
	"log"

	"github.com/rifqi142/indico-be/internal/models"
	"gorm.io/gorm"
)

func RunAutoMigration(db *gorm.DB) error {
	log.Println("Running auto migration...")

	err := db.AutoMigrate(
		&models.Voucher{},
	)

	if err != nil {
		return err
	}

	log.Println("Auto migration completed successfully")
	return nil
}
