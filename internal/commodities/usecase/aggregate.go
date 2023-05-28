package usecase

import (
	"context"
	"strings"

	"github.com/fidellr/fishery_api/internal/commodities"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type aggregateUsecase struct {
	aggregateRepo commodities.AggregateRepo
	context       context.Context
}

func NewAggregateUsecase(ctx context.Context, aggregateRepo commodities.AggregateRepo) commodities.AggregateUsecase {
	return &aggregateUsecase{
		aggregateRepo: aggregateRepo,
		context:       ctx,
	}
}

func (u *aggregateUsecase) GetMostCommodityRecords(ctx context.Context) ([]model.Commodity, error) {
	list, err := u.aggregateRepo.GetMostCommodityRecords(ctx)
	if err != nil {
		return nil, err
	}

	var mostRecords []model.Commodity

	counts := make(map[string]int)
	maxCount := 0

	for _, rec := range list {
		counts[rec.Komoditas]++
		if counts[rec.Komoditas] > maxCount {
			maxCount = counts[rec.Komoditas]
		}
	}

	for _, rec := range list {
		if counts[rec.Komoditas] == maxCount {
			mostRecords = append(mostRecords, rec)
		}
	}

	if len(mostRecords) == 0 {
		return nil, utils.ConstraintErrorf("No records found")
	}

	return mostRecords, nil
}

func (u *aggregateUsecase) GetByArea(ctx context.Context, city, province string) ([]model.Commodity, error) {
	list, err := u.aggregateRepo.GetByArea(ctx, strings.ToUpper(city), strings.ToUpper(province))
	if err != nil {
		return nil, err
	}

	return list, nil
}
