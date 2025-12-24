package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/rifqi142/indico-be/internal/dto"
	"github.com/rifqi142/indico-be/internal/models"
	"github.com/rifqi142/indico-be/internal/repository"
	"github.com/rifqi142/indico-be/internal/utils"
	"gorm.io/gorm"
)

type VoucherService interface {
	CreateVoucher(req dto.CreateVoucherRequest) (*dto.VoucherResponse, error)
	GetVoucherByID(id uint) (*dto.VoucherResponse, error)
	GetAllVouchers(query dto.VoucherListQuery) (*dto.VoucherListResponse, error)
	UpdateVoucher(id uint, req dto.UpdateVoucherRequest) (*dto.VoucherResponse, error)
	DeleteVoucher(id uint) error
	ImportFromCSV(reader io.Reader) (*dto.CSVUploadResponse, error)
	ExportToCSV() ([][]string, error)
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo: repo}
}

func (s *voucherService) CreateVoucher(req dto.CreateVoucherRequest) (*dto.VoucherResponse, error) {
	existing, _ := s.repo.FindByCode(req.Code)
	if existing != nil {
		return nil, errors.New("voucher code already exists")
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	voucher := &models.Voucher{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Discount:    req.Discount,
		MaxUsage:    req.MaxUsage,
		ValidFrom:   req.ValidFrom,
		ValidUntil:  req.ValidUntil,
		IsActive:    isActive,
	}

	if err := s.repo.Create(voucher); err != nil {
		return nil, err
	}

	return s.toVoucherResponse(voucher), nil
}

func (s *voucherService) GetVoucherByID(id uint) (*dto.VoucherResponse, error) {
	voucher, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		return nil, err
	}

	return s.toVoucherResponse(voucher), nil
}

func (s *voucherService) GetAllVouchers(query dto.VoucherListQuery) (*dto.VoucherListResponse, error) {
	vouchers, total, err := s.repo.FindAll(query)
	if err != nil {
		return nil, err
	}

	page := 1
	pageSize := 10
	if query.Page > 0 {
		page = query.Page
	}
	if query.PageSize > 0 {
		pageSize = query.PageSize
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	voucherResponses := make([]dto.VoucherResponse, len(vouchers))
	for i, voucher := range vouchers {
		voucherResponses[i] = *s.toVoucherResponse(&voucher)
	}

	return &dto.VoucherListResponse{
		Data: voucherResponses,
		Pagination: dto.PaginationMeta{
			CurrentPage: page,
			PageSize:    pageSize,
			TotalPages:  totalPages,
			TotalItems:  total,
		},
	}, nil
}

func (s *voucherService) UpdateVoucher(id uint, req dto.UpdateVoucherRequest) (*dto.VoucherResponse, error) {
	voucher, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		return nil, err
	}

	// Update fields if provided
	if req.Code != "" {
		existing, _ := s.repo.FindByCode(req.Code)
		if existing != nil && existing.ID != voucher.ID {
			return nil, errors.New("voucher code already exists")
		}
		voucher.Code = req.Code
	}
	if req.Name != "" {
		voucher.Name = req.Name
	}
	if req.Description != "" {
		voucher.Description = req.Description
	}
	if req.Discount > 0 {
		voucher.Discount = req.Discount
	}
	if req.MaxUsage > 0 {
		voucher.MaxUsage = req.MaxUsage
	}
	if !req.ValidFrom.IsZero() {
		voucher.ValidFrom = req.ValidFrom
	}
	if !req.ValidUntil.IsZero() {
		voucher.ValidUntil = req.ValidUntil
	}
	if req.IsActive != nil {
		voucher.IsActive = *req.IsActive
	}

	if err := s.repo.Update(voucher); err != nil {
		return nil, err
	}

	return s.toVoucherResponse(voucher), nil
}

func (s *voucherService) DeleteVoucher(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("voucher not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

func (s *voucherService) ImportFromCSV(reader io.Reader) (*dto.CSVUploadResponse, error) {
	csvReader := csv.NewReader(reader)
	
	// Read header
	header, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Validate header
	expectedHeaders := []string{"code", "name", "description", "discount", "max_usage", "valid_from", "valid_until", "is_active"}
	if !validateCSVHeader(header, expectedHeaders) {
		return nil, errors.New("invalid CSV header format")
	}

	var vouchers []models.Voucher
	rowNum := 1

	// Read data rows
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV row %d: %w", rowNum, err)
		}

		rowNum++

		voucher, err := s.parseCSVRow(record)
		if err != nil {
			continue
		}

		vouchers = append(vouchers, *voucher)
	}

	successCount, errors := s.repo.BulkCreate(vouchers)

	return &dto.CSVUploadResponse{
		SuccessCount: successCount,
		FailedCount:  len(vouchers) - successCount,
		Errors:       errors,
	}, nil
}

func (s *voucherService) ExportToCSV() ([][]string, error) {
	vouchers, err := s.repo.ExportAll()
	if err != nil {
		return nil, err
	}

	// Create CSV data
	data := [][]string{
		{"code", "name", "description", "discount", "max_usage", "used_count", "valid_from", "valid_until", "is_active", "created_at"},
	}

	for _, voucher := range vouchers {
		row := []string{
			voucher.Code,
			voucher.Name,
			voucher.Description,
			fmt.Sprintf("%.2f", voucher.Discount),
			strconv.Itoa(voucher.MaxUsage),
			strconv.Itoa(voucher.UsedCount),
			voucher.ValidFrom.Format("2006-01-02 15:04:05"),
			voucher.ValidUntil.Format("2006-01-02 15:04:05"),
			strconv.FormatBool(voucher.IsActive),
			voucher.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		data = append(data, row)
	}

	return data, nil
}

func (s *voucherService) parseCSVRow(record []string) (*models.Voucher, error) {
	if len(record) < 8 {
		return nil, errors.New("invalid number of columns")
	}

	discount, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid discount value")
	}

	maxUsage, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("invalid max_usage value")
	}

	validFrom, err := time.Parse("2006-01-02", strings.TrimSpace(record[5]))
	if err != nil {
		return nil, fmt.Errorf("invalid valid_from date")
	}

	validUntil, err := time.Parse("2006-01-02", strings.TrimSpace(record[6]))
	if err != nil {
		return nil, fmt.Errorf("invalid valid_until date")
	}

	isActive := true
	if len(record) > 7 {
		isActive, _ = strconv.ParseBool(record[7])
	}

	return &models.Voucher{
		Code:        strings.TrimSpace(record[0]),
		Name:        strings.TrimSpace(record[1]),
		Description: strings.TrimSpace(record[2]),
		Discount:    discount,
		MaxUsage:    maxUsage,
		ValidFrom:   validFrom,
		ValidUntil:  validUntil,
		IsActive:    isActive,
	}, nil
}

func (s *voucherService) toVoucherResponse(voucher *models.Voucher) *dto.VoucherResponse {
	return &dto.VoucherResponse{
		ID:          voucher.ID,
		Code:        voucher.Code,
		Name:        voucher.Name,
		Description: voucher.Description,
		Discount:    voucher.Discount,
		MaxUsage:    voucher.MaxUsage,
		UsedCount:   voucher.UsedCount,
		ValidFrom:   utils.NewReadableTime(voucher.ValidFrom),
		ValidUntil:  utils.NewReadableTime(voucher.ValidUntil),
		IsActive:    voucher.IsActive,
		CreatedAt:   utils.NewReadableTime(voucher.CreatedAt),
		UpdatedAt:   utils.NewReadableTime(voucher.UpdatedAt),
	}
}

func validateCSVHeader(header, expected []string) bool {
	if len(header) < len(expected) {
		return false
	}
	for i, h := range expected {
		if strings.ToLower(strings.TrimSpace(header[i])) != h {
			return false
		}
	}
	return true
}
