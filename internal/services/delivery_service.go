package services

import (
	"context"
	"fmt"
	"thanhnt208/delivery-service/external/client"
	"thanhnt208/delivery-service/internal/models"
	"thanhnt208/delivery-service/internal/repositories"
)

type DeliveryService interface {
	CreateDelivery(ctx context.Context, req *models.CreateDeliveryRequest) (*models.DeliveryResponse, error)
	CalculateDistance(ctx context.Context, from, to string) (*models.DistanceResponse, error)
}

type deliveryService struct {
	repo         repositories.DeliveryRepository
	mapboxClient *client.MapboxClient
}

func NewDeliveryService(repo repositories.DeliveryRepository, mapboxClient *client.MapboxClient) DeliveryService {
	return &deliveryService{
		repo:         repo,
		mapboxClient: mapboxClient,
	}
}

func (s *deliveryService) CalculateDistance(ctx context.Context, from, to string) (*models.DistanceResponse, error) {
	fromCoords, err := s.mapboxClient.GeocodeAddress(from)
	if err != nil {
		return nil, fmt.Errorf("failed to geocode 'from' address: %w", err)
	}

	toCoords, err := s.mapboxClient.GeocodeAddress(to)
	if err != nil {
		return nil, fmt.Errorf("failed to geocode 'to' address: %w", err)
	}

	distance, duration, geometry, err := s.mapboxClient.GetDirections(fromCoords, toCoords)
	if err != nil {
		return nil, fmt.Errorf("failed to get directions: %w", err)
	}

	return &models.DistanceResponse{
		Distance:     distance / 1000,                      // Convert to kilometers
		Duration:     duration / 60,                        // Convert to minutes
		Fee:          int64(float64(distance/1000) * 5000), // 5000 per km
		FromCoords:   fromCoords,
		ToCoords:     toCoords,
		GeometryLine: geometry,
	}, nil
}

func (s *deliveryService) CreateDelivery(ctx context.Context, req *models.CreateDeliveryRequest) (*models.DeliveryResponse, error) {
	distanceResult, err := s.CalculateDistance(ctx, req.RestaurantAddress, req.ShippingAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate distance: %w", err)
	}

	shipper, err := s.repo.GetAvailableShipper(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get available shipper: %w", err)
	}

	delivery := &models.Delivery{
		OrderID:           req.OrderID,
		ShipperID:         shipper.ID,
		RestaurantAddress: req.RestaurantAddress,
		ShippingAddress:   req.ShippingAddress,
		Distance:          distanceResult.Distance,
		Duration:          distanceResult.Duration,
		Fee:               distanceResult.Fee,
		FromCoords:        distanceResult.FromCoords,
		ToCoords:          distanceResult.ToCoords,
		GeometryLine:      distanceResult.GeometryLine,
		Status:            "assigned",
	}

	deliveryID, err := s.repo.CreateDelivery(ctx, delivery)
	if err != nil {
		return nil, fmt.Errorf("failed to create delivery: %w", err)
	}

	err = s.repo.AssignDelivery(ctx, req.OrderID, shipper.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to assign shipper: %w", err)
	}

	err = s.repo.UpdateShipperStatus(ctx, shipper.ID, "delivering")
	if err != nil {
		return nil, fmt.Errorf("failed to update shipper status: %w", err)
	}

	shipper.Status = "delivering"

	return &models.DeliveryResponse{
		DeliveryID:   deliveryID,
		OrderID:      req.OrderID,
		Distance:     distanceResult.Distance,
		Duration:     distanceResult.Duration,
		Fee:          distanceResult.Fee,
		FromCoords:   distanceResult.FromCoords,
		ToCoords:     distanceResult.ToCoords,
		GeometryLine: distanceResult.GeometryLine,
		Status:       "assigned",
		Shipper:      *shipper,
	}, nil
}
