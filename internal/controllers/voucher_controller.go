package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rifqi142/indico-be/internal/dto"
	"github.com/rifqi142/indico-be/internal/services"
	"github.com/rifqi142/indico-be/internal/utils"
)

type VoucherController struct {
	voucherService services.VoucherService
}

func NewVoucherController(voucherService services.VoucherService) *VoucherController {
	return &VoucherController{voucherService: voucherService}
}

func (ctrl *VoucherController) GetAllVouchers(c *gin.Context) {
	var query dto.VoucherListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.BadRequestResponse(c, "Invalid query parameters", err.Error())
		return
	}

	result, err := ctrl.voucherService.GetAllVouchers(query)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get vouchers", err.Error())
		return
	}

	utils.SuccessResponse(c, "Vouchers retrieved successfully", result)
}

func (ctrl *VoucherController) GetVoucherByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid voucher ID", err.Error())
		return
	}

	result, err := ctrl.voucherService.GetVoucherByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, "Voucher retrieved successfully", result)
}

func (ctrl *VoucherController) CreateVoucher(c *gin.Context) {
	var req dto.CreateVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	result, err := ctrl.voucherService.CreateVoucher(req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), nil)
		return
	}

	utils.CreatedResponse(c, "Voucher created successfully", result)
}

func (ctrl *VoucherController) UpdateVoucher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid voucher ID", err.Error())
		return
	}

	var req dto.UpdateVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	result, err := ctrl.voucherService.UpdateVoucher(uint(id), req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "Voucher updated successfully", result)
}

func (ctrl *VoucherController) DeleteVoucher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid voucher ID", err.Error())
		return
	}

	if err := ctrl.voucherService.DeleteVoucher(uint(id)); err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, "Voucher deleted successfully", nil)
}

func (ctrl *VoucherController) UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.BadRequestResponse(c, "File is required", err.Error())
		return
	}

	if file.Header.Get("Content-Type") != "text/csv" && !isCSVFile(file.Filename) {
		utils.BadRequestResponse(c, "Only CSV files are allowed", nil)
		return
	}

	src, err := file.Open()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to open file", err.Error())
		return
	}
	defer src.Close()

	// Process CSV
	result, err := ctrl.voucherService.ImportFromCSV(src)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "CSV uploaded successfully", result)
}

func (ctrl *VoucherController) ExportCSV(c *gin.Context) {
	data, err := ctrl.voucherService.ExportToCSV()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to export vouchers", err.Error())
		return
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	if err := writer.WriteAll(data); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to write CSV", err.Error())
		return
	}

	writer.Flush()

	// Set headers for file download
	filename := fmt.Sprintf("vouchers_export_%s.csv", time.Now().Format("20060102_150405"))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", buf.Bytes())
}

func isCSVFile(filename string) bool {
	return len(filename) > 4 && filename[len(filename)-4:] == ".csv"
}
