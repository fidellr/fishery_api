package http

import (
	"context"
	"encoding/json"

	"github.com/fidellr/fishery_api/internal/commodities"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type Handler struct {
	aggregateService commodities.AggregateUsecase
	commodityService commodities.CommodityUsecase
}

func NewCommoditiesHandler(commodityService commodities.CommodityUsecase, aggregateService commodities.AggregateUsecase) *Handler {
	handler := &Handler{commodityService: commodityService, aggregateService: aggregateService}
	return handler
}

func (h *Handler) AddRecords(ctx context.Context, records []model.Commodity) error {
	err := h.commodityService.AddRecords(ctx, records)
	if err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) GetAllByCommodity(ctx context.Context, commodity string) (string, error) {
	list, err := h.commodityService.GetAllByCommodity(ctx, commodity)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}

func (h *Handler) GetByID(ctx context.Context, uuid string) (string, error) {
	list, err := h.commodityService.GetByID(ctx, uuid)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}

func (h *Handler) UpdateRecords(ctx context.Context, payloads []model.Commodity) error {
	if err := h.commodityService.UpdateRecords(ctx, payloads); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}

func (h *Handler) DeleteRecords(ctx context.Context, uuid []string) error {
	if err := h.commodityService.DeleteRecords(ctx, uuid); err != nil {
		return utils.ConstraintErrorf("%s", err.Error())
	}

	return nil
}
