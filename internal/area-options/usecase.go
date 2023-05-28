package areaoptions

import (
	"context"

	"github.com/fidellr/fishery_api/internal/model"
)

type AreaOptionUsecase interface {
	AddRecords(ctx context.Context, records []model.AreaOption) error
	UpdateRecords(ctx context.Context, payloads []model.AreaOption) error
	DeleteRecords(ctx context.Context, records []model.AreaOption) error
}
