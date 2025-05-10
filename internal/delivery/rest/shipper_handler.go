package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"thanhnt208/delivery-service/internal/models"
	"thanhnt208/delivery-service/internal/services"

	"github.com/gin-gonic/gin"
)

type ShipperHandler struct {
	service services.ShipperService
}

func NewShipperHandler(service services.ShipperService) *ShipperHandler {
	return &ShipperHandler{service: service}
}

// POST /shippers
func (h *ShipperHandler) CreateShipper(c *gin.Context) {
	var req models.ShipperRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateShipper(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET /shippers/:id
func (h *ShipperHandler) GetShipperByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, err := h.service.GetShipperByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET /shippers?limit=10&offset=0
func (h *ShipperHandler) ListShippers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	results, err := h.service.ListShippers(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *DeliveryHandler) GetDeliveryByOrderID(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
	if err != nil {
		// If the orderId is invalid (cannot convert to int64)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid orderId format"})
		return
	}

	deliveries, err := h.service.GetDeliveryByOrderID(c.Request.Context(), orderId)
	if err != nil {
		// If there is an error retrieving deliveries from the service
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("delivery for orderId %d not found: %v", orderId, err),
		})
		return
	}

	// Return the deliveries if everything is fine
	c.JSON(http.StatusOK, deliveries)
}
