package commodities

import (
	"context"
	"time"

	"github.com/fidellr/fishery_api/internal/model"
)

type AggregateUsecase interface {
	GetAllByRange(ctx context.Context, price, size string, date time.Time) ([]*model.Commodity, error)
	GetMaxPrice(ctx context.Context, week int32, commodity string) ([]*model.Commodity, error)
	GetMostRecordsByCommodity(ctx context.Context, commodity string) ([]*model.Commodity, error)
	GetByArea(city, province string) ([]*model.Commodity, error)
}

type CommodityUsecase interface {
	// AddRecords(ctx context.Context, records []*model.Commodity) error
	GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error)
	// GetByID(ctx context.Context, uuid string) (*model.Commodity, error)
	// UpdateRecords(ctx context.Context, uuid []string) error
	// DeleteRecords(ctx context.Context, uuid []string) error
}
