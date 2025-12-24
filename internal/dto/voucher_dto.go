package dto

import (
	"time"

	"github.com/rifqi142/indico-be/internal/utils"
)

type CreateVoucherRequest struct {
	Code        string    `json:"code" binding:"required,min=3,max=50"`
	Name        string    `json:"name" binding:"required,min=3,max=255"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount" binding:"required,min=0,max=100"`
	MaxUsage    int       `json:"max_usage" binding:"required,min=1"`
	ValidFrom   time.Time `json:"valid_from" binding:"required"`
	ValidUntil  time.Time `json:"valid_until" binding:"required,gtfield=ValidFrom"`
	IsActive    *bool     `json:"is_active"`
}

type UpdateVoucherRequest struct {
	Code 	  string    `json:"code" binding:"omitempty,min=3,max=50"`
	Name        string    `json:"name" binding:"omitempty,min=3,max=255"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount" binding:"omitempty,min=0,max=100"`
	MaxUsage    int       `json:"max_usage" binding:"omitempty,min=1"`
	ValidFrom   time.Time `json:"valid_from"`
	ValidUntil  time.Time `json:"valid_until"`
	IsActive    *bool     `json:"is_active"`
}

type VoucherResponse struct {
	ID          uint                `json:"id"`
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Discount    float64             `json:"discount"`
	MaxUsage    int                 `json:"max_usage"`
	UsedCount   int                 `json:"used_count"`
	ValidFrom   utils.ReadableTime  `json:"valid_from"`
	ValidUntil  utils.ReadableTime  `json:"valid_until"`
	IsActive    bool                `json:"is_active"`
	CreatedAt   utils.ReadableTime  `json:"created_at"`
	UpdatedAt   utils.ReadableTime  `json:"updated_at"`
}

type VoucherListQuery struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	Search   string `form:"search"`
	SortBy   string `form:"sort_by" binding:"omitempty,oneof=id code name discount created_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	IsActive *bool  `form:"is_active"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
}

type VoucherListResponse struct {
	Data       []VoucherResponse `json:"data"`
	Pagination PaginationMeta    `json:"pagination"`
}

type CSVUploadResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedCount  int      `json:"failed_count"`
	Errors       []string `json:"errors,omitempty"`
}
