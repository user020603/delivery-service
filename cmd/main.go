package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"thanhnt208/delivery-service/config"
	"thanhnt208/delivery-service/external/client"
	"thanhnt208/delivery-service/internal/delivery/rest"
	"thanhnt208/delivery-service/internal/repositories"
	"thanhnt208/delivery-service/internal/services"
)

func main() {
	_ = godotenv.Load()

	db, err := config.ConnectPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userClient := &client.UserClient{}
	shipperRepo := repositories.NewShipperRepository(db)
	shipperService := services.NewShipperService(shipperRepo, userClient)
	shipperHandler := rest.NewShipperHandler(shipperService)

	r := gin.Default()

	r.POST("/shippers", shipperHandler.CreateShipper)
	r.GET("/shippers/:id", shipperHandler.GetShipperByID)
	r.GET("/shippers", shipperHandler.ListShippers)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
