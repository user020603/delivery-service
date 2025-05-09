package repositories

import (
	"context"
	"thanhnt208/delivery-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type ShipperRepository interface {
	Create(ctx context.Context, shipper *models.ShipperRequest) (*models.ShipperResponse, error)
	GetByID(ctx context.Context, id int64) (*models.ShipperResponse, error)
	GetShippers(ctx context.Context, limit, offset int) ([]*models.ShipperResponse, error)
}

type shipperRepository struct {
	db *sqlx.DB
}

func NewShipperRepository(db *sqlx.DB) ShipperRepository {
	return &shipperRepository{db: db}
}

func (r *shipperRepository) Create(ctx context.Context, shipper *models.ShipperRequest) (*models.ShipperResponse, error) {
	var id int64
	query := `
		INSERT INTO shippers (email, name, gender, phone, vehicle_type, vehicle_plate, total_deliveries, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	err := r.db.QueryRowContext(
		ctx, query,
		shipper.Email,
		shipper.Name,
		shipper.Gender,
		shipper.Phone,
		shipper.VehicleType,
		shipper.VehiclePlate,
	).Scan(&id)
	return &models.ShipperResponse{
		ID:              id,
		Email:           shipper.Email,
		Name:            shipper.Name,
		Gender:          shipper.Gender,
		Phone:           shipper.Phone,
		Role:            "shipper",
		VehicleType:     shipper.VehicleType,
		VehiclePlate:    shipper.VehiclePlate,
		TotalDeliveries: 0,
		Status:          "available",
	}, err
}

func (r *shipperRepository) GetByID(ctx context.Context, id int64) (*models.ShipperResponse, error) {
	var shipper models.ShipperResponse
	query := `SELECT * FROM shippers WHERE id = $1`
	err := r.db.GetContext(ctx, &shipper, query, id)
	if err != nil {
		return nil, err
	}
	return &shipper, nil
}

func (r *shipperRepository) GetShippers(ctx context.Context, limit, offset int) ([]*models.ShipperResponse, error) {
	var shippers []*models.ShipperResponse
	query := `SELECT * FROM shippers LIMIT $1 OFFSET $2`
	err := r.db.SelectContext(ctx, &shippers, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return shippers, nil
}
