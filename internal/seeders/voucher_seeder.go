package seeders

import (
	"log"
	"time"

	"github.com/rifqi142/indico-be/internal/models"
	"gorm.io/gorm"
)

func SeedVouchers(db *gorm.DB) {
	log.Println("Seeding vouchers...")

	vouchers := []models.Voucher{
		{
			Code:        "WELCOME2025",
			Name:        "Welcome Bonus 2025",
			Description: "Special discount for new customers in 2025",
			Discount:    25.00,
			MaxUsage:    100,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "NEWYEAR50",
			Name:        "New Year Flash Sale",
			Description: "Limited time 50% discount for New Year celebration",
			Discount:    50.00,
			MaxUsage:    50,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 1, 7, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "VALENTINE20",
			Name:        "Valentine Special",
			Description: "Show love with 20% discount",
			Discount:    20.00,
			MaxUsage:    200,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 2, 10, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 2, 14, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "SPRING15",
			Name:        "Spring Sale",
			Description: "Fresh start with 15% off",
			Discount:    15.00,
			MaxUsage:    150,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 5, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "SUMMER30",
			Name:        "Summer Vibes",
			Description: "Hot deals with 30% discount",
			Discount:    30.00,
			MaxUsage:    100,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 8, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "BACKTOSCHOOL",
			Name:        "Back to School",
			Description: "Student discount 25% off",
			Discount:    25.00,
			MaxUsage:    300,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 10, 1, 6, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "OCTOBER10",
			Name:        "October Fest",
			Description: "Celebrate with 10% discount",
			Discount:    10.00,
			MaxUsage:    500,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 11, 1, 6, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "BLACKFRIDAY",
			Name:        "Black Friday Mega Sale",
			Description: "Biggest discount of the year - 60% off",
			Discount:    60.00,
			MaxUsage:    200,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 11, 28, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 11, 30, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "CYBERMONDAY",
			Name:        "Cyber Monday Special",
			Description: "Online exclusive 45% discount",
			Discount:    45.00,
			MaxUsage:    250,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 2, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "CHRISTMAS35",
			Name:        "Christmas Gift",
			Description: "Holiday season special 35% off",
			Discount:    35.00,
			MaxUsage:    400,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 12, 15, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 25, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "VIPGOLD",
			Name:        "VIP Gold Member",
			Description: "Exclusive VIP discount 40%",
			Discount:    40.00,
			MaxUsage:    1000,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "FIRSTBUY",
			Name:        "First Purchase",
			Description: "First time buyer gets 30% off",
			Discount:    30.00,
			MaxUsage:    500,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "LOYAL100",
			Name:        "Loyalty Reward",
			Description: "Thank you for being loyal - 20% off",
			Discount:    20.00,
			MaxUsage:    1000,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "FLASH5MIN",
			Name:        "Flash 5 Minutes",
			Description: "Ultra limited 70% discount - only 20 uses",
			Discount:    70.00,
			MaxUsage:    20,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 1, 15, 12, 5, 0, 0, time.UTC),
			IsActive:    true,
		},
		{
			Code:        "WEEKEND15",
			Name:        "Weekend Special",
			Description: "Every weekend get 15% off",
			Discount:    15.00,
			MaxUsage:    1000,
			UsedCount:   0,
			ValidFrom:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			ValidUntil:  time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			IsActive:    true,
		},
	}

	for _, voucher := range vouchers {
		err := db.Where(models.Voucher{Code: voucher.Code}).FirstOrCreate(&voucher).Error
		if err != nil {
			log.Printf("Failed to seed voucher %s: %v\n", voucher.Code, err)
		}
	}

	log.Println("Vouchers seeding completed successfully")
}
