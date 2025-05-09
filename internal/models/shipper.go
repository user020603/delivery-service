package models

type Shipper struct {
	ID              int64  `db:"id"`
	Password        string `db:"password"`
	Email           string `db:"email"`
	Name            string `db:"name"`
	Gender          string `db:"gender"`
	Phone           string `db:"phone"`
	VehicleType     string `db:"vehicle_type"`
	VehiclePlate    string `db:"vehicle_plate"`
	TotalDeliveries int    `db:"total_deliveries"`
	Status          string `db:"status"`
}

type ShipperRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	Name         string `json:"name" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	VehicleType  string `json:"vehicleType" binding:"required"`
	VehiclePlate string `json:"vehiclePlate" binding:"required"`
}

type ShipperResponse struct {
	ID              int64  `json:"userId"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Phone           string `json:"phone"`
	Role            string `json:"role"`
	VehicleType     string `json:"vehicleType"`
	VehiclePlate    string `json:"vehiclePlate"`
	TotalDeliveries int    `json:"totalDeliveries"`
	Status          string `json:"status"`
}
