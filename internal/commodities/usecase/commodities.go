package usecase

import (
	"context"

	"github.com/fidellr/fishery_api/internal/commodities"
	"github.com/fidellr/fishery_api/internal/model"
)

type commoditiesUsecase struct {
	commoditiesRepo commodities.CommodityRepo
	context         context.Context
}

func NewCommoditiesUsecase(ctx context.Context, commoditiesRepo commodities.CommodityRepo) commodities.CommodityUsecase {
	return &commoditiesUsecase{
		commoditiesRepo: commoditiesRepo,
		context:         ctx,
	}
}

func (u *commoditiesUsecase) GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error) {
	list, err := u.commoditiesRepo.GetAllByCommodity(ctx, commodity)
	if err != nil {
		return nil, err
	}

	return list, nil
}
