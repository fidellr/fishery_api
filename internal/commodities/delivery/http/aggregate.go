package http

import (
	"context"
	"encoding/json"

	"github.com/fidellr/fishery_api/internal/utils"
)

func (h *Handler) GetMostCommodityRecords(ctx context.Context) (string, error) {
	list, err := h.aggregateService.GetMostCommodityRecords(ctx)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}

func (h *Handler) GetByArea(ctx context.Context, city, province string) (string, error) {
	list, err := h.aggregateService.GetByArea(ctx, city, province)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}

func (h *Handler) GetMaxPrice(ctx context.Context, week int, commodity string) (string, error) {
	list, err := h.aggregateService.GetMaxPrice(ctx, week, commodity)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}

func (h *Handler) GetAllByRange(ctx context.Context, price, size, date string) (string, error) {
	list, err := h.aggregateService.GetAllByRange(ctx, price, size, date)
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return "", utils.ConstraintErrorf("%s", err.Error())
	}

	return string(jsonByte), nil
}
