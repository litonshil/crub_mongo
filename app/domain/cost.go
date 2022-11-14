package domain

import (
	"crud_mongo/app/models"
	"crud_mongo/app/serializers"
)

type ICostRepo interface {
	InsertDailyCost(req *models.Cost) error
	GetDailyCost() ([]models.Cost, error)
}

type ICostSvc interface {
	InsertDailyCost(payload *serializers.CostPayload) error
	GetDailyCost() ([]models.Cost, error)
	// GetCostSummary() ([]models.Cost, error)
}
