package commodities

import (
	"context"

	"github.com/fidellr/fishery_api/internal/model"
)

type AggregateUsecase interface {
	GetAllByRange(ctx context.Context, price, size, date string) ([]model.Commodity, error)
	GetMaxPrice(ctx context.Context, week int, commodity string) ([]model.Commodity, error)
	GetMostCommodityRecords(ctx context.Context) ([]model.Commodity, error)
	GetByArea(ctx context.Context, city, province string) ([]model.Commodity, error)
}

type CommodityUsecase interface {
	AddRecords(ctx context.Context, records []model.Commodity) error
	GetAllCommodity(ctx context.Context) ([]model.Commodity, error)
	GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error)
	GetByID(ctx context.Context, uuid string) ([]model.Commodity, error)
	UpdateRecords(ctx context.Context, payloads []model.Commodity) error
	DeleteRecords(ctx context.Context, uuid []string) error
}
