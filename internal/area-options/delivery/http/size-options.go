package http

import (
	"context"

	areaoptions "github.com/fidellr/fishery_api/internal/area-options"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type Handler struct {
	areaOptService areaoptions.AreaOptionUsecase
}

func NewAreaOptsHandler(areaOptService areaoptions.AreaOptionUsecase) *Handler {
	handler := &Handler{areaOptService}
	return handler
}

func (h *Handler) AddRecords(ctx context.Context, records []model.AreaOption) error {
	err := h.areaOptService.AddRecords(ctx, records)
	if err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) UpdateRecords(ctx context.Context, payloads []model.AreaOption) error {
	if err := h.areaOptService.UpdateRecords(ctx, payloads); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) DeleteRecords(ctx context.Context, records []model.AreaOption) error {
	if err := h.areaOptService.DeleteRecords(ctx, records); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}
