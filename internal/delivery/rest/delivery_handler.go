package rest

import (
	"net/http"
	"thanhnt208/delivery-service/internal/models"
	"thanhnt208/delivery-service/internal/services"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	service services.DeliveryService
}

func NewDeliveryHandler(service services.DeliveryService) (*DeliveryHandler, error) {
	return &DeliveryHandler{service: service}, nil
}

// POST /delivery
func (h *DeliveryHandler) CreateDelivery(c *gin.Context) {
	var req models.CreateDeliveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateDelivery(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
