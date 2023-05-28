package sizeoptions

import (
	"context"

	"github.com/fidellr/fishery_api/internal/model"
)

type SizeOptionRepo interface {
	AddRecords(ctx context.Context, records []model.SizeOption) error
	UpdateRecords(ctx context.Context, payloads []model.SizeOption) error
	DeleteRecords(ctx context.Context, records []model.SizeOption) error
}
