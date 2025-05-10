package routes

import (
	"net/http"
	"thanhnt208/delivery-service/api/middlewares"
	"thanhnt208/delivery-service/internal/delivery/rest"
	"thanhnt208/delivery-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(shipperHandler *rest.ShipperHandler, deliveryHandler *rest.DeliveryHandler) *gin.Engine {
	router := gin.Default()

	logger := logger.NewLogger("info")

	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.LoggingMiddleware(logger))
	router.Use(middlewares.RecoveryMiddleware(logger))
	// router.Use(middlewares.AuthAdminMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"\nversion": "1.0.0\n",
			"\nstatus":  "UP\n",
		})
	})

	v1 := router.Group("/api/v1")
	{
		setupShipperRoutes(v1, shipperHandler)
		setupDeliveryRouters(v1, deliveryHandler)
	}

	return router
}

func setupShipperRoutes(rg *gin.RouterGroup, handler *rest.ShipperHandler) {
	shipper := rg.Group("/shippers")
	{
		shipper.POST("/", handler.CreateShipper)
		shipper.GET("/:id", handler.GetShipperByID)
		shipper.GET("/", handler.ListShippers)
	}
}

func setupDeliveryRouters(rg *gin.RouterGroup, handler *rest.DeliveryHandler) {
	delivery := rg.Group("/deliveries")
	{
		delivery.POST("/", handler.CreateDelivery)
		delivery.PUT("/:deliveryId/status", handler.UpdateDeliveryStatus)
		delivery.GET("/shipper/:shipperId", handler.GetDeliveriesByShipperID)
	}
}
