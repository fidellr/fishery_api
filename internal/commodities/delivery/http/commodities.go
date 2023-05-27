package http

import (
	"context"
	"encoding/json"

	"github.com/fidellr/fishery_api/internal/commodities"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type Handler struct {
	commodityService commodities.CommodityUsecase
}

type Response struct {
	res []*model.Commodity
	err error
}

func NewCommoditiesHandler(commodityService commodities.CommodityUsecase) *Handler {
	handler := &Handler{commodityService}
	return handler
}

func (h *Handler) GetAllByCommodity(ctx context.Context, commodity string) ([]byte, error) {
	list, err := h.commodityService.GetAllByCommodity(ctx, commodity)
	if err != nil {
		return nil, utils.ConstraintErrorf("%s", err.Error())
	}

	jsonByte, err := json.Marshal(list)
	if err != nil {
		return nil, utils.ConstraintErrorf("%s", err.Error())
	}

	return jsonByte, nil
}
