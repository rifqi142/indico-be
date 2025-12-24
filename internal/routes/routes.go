package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqi142/indico-be/internal/controllers"
	"github.com/rifqi142/indico-be/internal/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
	voucherController *controllers.VoucherController,
	jwtSecret string,
) {
	router.Use(middleware.CORSMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	router.POST("/login", authController.Login)

	api := router.Group("/")
	api.Use(middleware.AuthMiddleware(jwtSecret))
	{
		vouchers := api.Group("/vouchers")
		{
			vouchers.GET("", voucherController.GetAllVouchers)
			vouchers.GET("/get-by-id/:id", voucherController.GetVoucherByID)
			vouchers.POST("", voucherController.CreateVoucher)
			vouchers.PUT("/:id", voucherController.UpdateVoucher)
			vouchers.DELETE("/:id", voucherController.DeleteVoucher)
			
			// CSV operations
			vouchers.POST("/upload-csv", voucherController.UploadCSV)
			vouchers.GET("/export", voucherController.ExportCSV)
		}
	}
}
