package rest

import (
	"net/http"
	"strconv"
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

func (h *DeliveryHandler) UpdateDeliveryStatus(c *gin.Context) {
	deliveryIdStr := c.Param("deliveryId")
	deliveryId, err := strconv.ParseInt(deliveryIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deliveryId"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	if err := h.service.UpdateDeliveryStatus(c.Request.Context(), deliveryId, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery status updated successfully"})
}

func (h *DeliveryHandler) GetDeliveriesByShipperID(c *gin.Context) {
	shipperIdStr := c.Param("shipperId")
	shipperId, err := strconv.ParseInt(shipperIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shipperId"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	deliveries, err := h.service.GetDeliveriesByShipperID(c.Request.Context(), shipperId, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveries)
}
