package repository

import (
	"fmt"
	"strings"

	"github.com/rifqi142/indico-be/internal/dto"
	"github.com/rifqi142/indico-be/internal/models"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	Create(voucher *models.Voucher) error
	FindByID(id uint) (*models.Voucher, error)
	FindByCode(code string) (*models.Voucher, error)
	FindAll(query dto.VoucherListQuery) ([]models.Voucher, int64, error)
	Update(voucher *models.Voucher) error
	Delete(id uint) error
	BulkCreate(vouchers []models.Voucher) (int, []string)
	ExportAll() ([]models.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

func (r *voucherRepository) Create(voucher *models.Voucher) error {
	return r.db.Create(voucher).Error
}

func (r *voucherRepository) FindByID(id uint) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.First(&voucher, id).Error
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) FindByCode(code string) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.Where("code = ?", code).First(&voucher).Error
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) FindAll(query dto.VoucherListQuery) ([]models.Voucher, int64, error) {
	var vouchers []models.Voucher
	var total int64

	db := r.db.Model(&models.Voucher{})

	// Apply filter conditions
	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(code) LIKE ? OR LOWER(name) LIKE ? OR LOWER(description) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if query.IsActive != nil {
		db = db.Where("is_active = ?", *query.IsActive)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	sortOrder := "asc"
	if query.SortBy != "" {
		sortBy = query.SortBy
	}
	if query.SortOrder != "" {
		sortOrder = query.SortOrder
	}
	db = db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	page := 1
	pageSize := 10
	if query.Page > 0 {
		page = query.Page
	}
	if query.PageSize > 0 {
		pageSize = query.PageSize
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	if err := db.Find(&vouchers).Error; err != nil {
		return nil, 0, err
	}

	return vouchers, total, nil
}

func (r *voucherRepository) Update(voucher *models.Voucher) error {
	return r.db.Save(voucher).Error
}

func (r *voucherRepository) Delete(id uint) error {
	return r.db.Delete(&models.Voucher{}, id).Error
}

func (r *voucherRepository) BulkCreate(vouchers []models.Voucher) (int, []string) {
	successCount := 0
	var errors []string

	for i, voucher := range vouchers {
		if err := r.db.Create(&voucher).Error; err != nil {
			errors = append(errors, fmt.Sprintf("Row %d: %s", i+1, err.Error()))
		} else {
			successCount++
		}
	}

	return successCount, errors
}

func (r *voucherRepository) ExportAll() ([]models.Voucher, error) {
	var vouchers []models.Voucher
	err := r.db.Order("created_at desc").Find(&vouchers).Error
	return vouchers, err
}
