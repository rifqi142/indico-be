package models

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"uniqueIndex;not null;size:50" json:"code"`
	Name        string         `gorm:"not null;size:255" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Discount    float64        `gorm:"not null" json:"discount"`
	MaxUsage    int            `gorm:"not null;default:1" json:"max_usage"`
	UsedCount   int            `gorm:"default:0" json:"used_count"`
	ValidFrom   time.Time      `gorm:"not null" json:"valid_from"`
	ValidUntil  time.Time      `gorm:"not null" json:"valid_until"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Voucher) TableName() string {
	return "vouchers"
}

func (v *Voucher) IsValid() bool {
	now := time.Now()
	return v.IsActive &&
		v.UsedCount < v.MaxUsage &&
		now.After(v.ValidFrom) &&
		now.Before(v.ValidUntil)
}

func (v *Voucher) CanBeUsed() bool {
	return v.IsValid() && v.UsedCount < v.MaxUsage
}

func (v *Voucher) IncrementUsage() {
	v.UsedCount++
}
