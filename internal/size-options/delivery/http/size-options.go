package http

import (
	"context"

	"github.com/fidellr/fishery_api/internal/model"
	sizeoptions "github.com/fidellr/fishery_api/internal/size-options"
	"github.com/fidellr/fishery_api/internal/utils"
)

type Handler struct {
	sizeOptService sizeoptions.SizeOptionUsecase
}

func NewSizeOptsHandler(sizeOptService sizeoptions.SizeOptionUsecase) *Handler {
	handler := &Handler{sizeOptService}
	return handler
}

func (h *Handler) AddRecords(ctx context.Context, records []model.SizeOption) error {
	err := h.sizeOptService.AddRecords(ctx, records)
	if err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) UpdateRecords(ctx context.Context, payloads []model.SizeOption) error {
	if err := h.sizeOptService.UpdateRecords(ctx, payloads); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) DeleteRecords(ctx context.Context, records []model.SizeOption) error {
	if err := h.sizeOptService.DeleteRecords(ctx, records); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}
