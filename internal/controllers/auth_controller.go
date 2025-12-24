package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqi142/indico-be/internal/dto"
	"github.com/rifqi142/indico-be/internal/services"
	"github.com/rifqi142/indico-be/internal/utils"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	result, err := ctrl.authService.Login(req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token", err.Error())
		return
	}

	utils.SuccessResponse(c, "Login successful", result)
}
