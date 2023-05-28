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
