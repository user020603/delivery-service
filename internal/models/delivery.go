package models

import (
	"time"
)

type CreateDeliveryRequest struct {
	OrderID           int64  `json:"orderId"`
	RestaurantAddress string `json:"restaurantAddress"`
	ShippingAddress   string `json:"shippingAddress"`
}

type DeliveryResponse struct {
	DeliveryID   int64           `json:"deliveryId"`
	OrderID      int64           `json:"orderId"`
	Distance     float64         `json:"distance"`
	Duration     float64         `json:"duration"`
	Fee          int64           `json:"fee"`
	FromCoords   []float64       `json:"fromCoords"`
	ToCoords     []float64       `json:"toCoords"`
	GeometryLine string          `json:"geometryLine"`
	Status       string          `json:"status"`
	Shipper      ShipperResponse `json:"shipper"`
}

type Delivery struct {
	ID                int64     `db:"id"`
	OrderID           int64     `db:"order_id"`
	ShipperID         int64     `db:"shipper_id"`
	RestaurantAddress string    `db:"restaurant_address"`
	ShippingAddress   string    `db:"shipping_address"`
	Distance          float64   `db:"distance"`
	Duration          float64   `db:"duration"`
	Fee               int64     `db:"fee"`
	FromCoords        []float64 `db:"from_coords"`
	ToCoords          []float64 `db:"to_coords"`
	GeometryLine      string    `db:"geometry_line"`
	Status            string    `db:"status"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type LocationRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type DistanceResponse struct {
	Distance     float64   `json:"distance"`
	Duration     float64   `json:"duration"`
	Fee          int64     `json:"fee"`
	FromCoords   []float64 `json:"fromCoords"`
	ToCoords     []float64 `json:"toCoords"`
	GeometryLine string    `json:"geometryLine,omitempty"`
}
