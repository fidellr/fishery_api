package usecase

import (
	"context"
	"strconv"
	"strings"
	"time"

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

func (u *aggregateUsecase) GetMaxPrice(ctx context.Context, week int, commodity string) ([]model.Commodity, error) {
	list, err := u.aggregateRepo.GetMaxPrice(ctx, week, commodity)
	if err != nil {
		return nil, err
	}

	var commodities []model.Commodity

	for _, rec := range list {
		timestampInt, err := strconv.ParseInt(rec.Timestamp, 10, 64)
		if err != nil {
			return nil, err
		}

		timestamp := time.Unix(0, timestampInt*int64(time.Millisecond))

		if timestamp.Before(time.Now().AddDate(0, 0, -(week * 7))) {
			commodities = append(commodities, model.Commodity{
				UUID:         rec.Timestamp,
				Komoditas:    rec.Komoditas,
				AreaProvinsi: rec.AreaProvinsi,
				AreaKota:     rec.AreaKota,
				Size:         rec.Size,
				Price:        rec.Price,
				TglParsed:    rec.TglParsed,
				Timestamp:    rec.Timestamp,
			})
		}
	}

	counts := make(map[string]int)
	maxCount := 0
	var maxPriceComm []model.Commodity
	for _, rec := range commodities {
		priceInt, err := strconv.ParseInt(rec.Price, 10, 64)
		if err != nil {
			return nil, err
		}

		counts[rec.Komoditas] = int(priceInt)
		if counts[rec.Komoditas] > maxCount {
			maxCount = counts[rec.Komoditas]
		}
	}

	for _, rec := range commodities {
		priceInt, err := strconv.ParseInt(rec.Price, 10, 64)
		if err != nil {
			return nil, err
		}

		if int(priceInt) == maxCount {
			maxPriceComm = append(maxPriceComm, rec)
		}
	}

	return maxPriceComm, nil
}

func (u *aggregateUsecase) GetAllByRange(ctx context.Context, price, size, date string) ([]model.Commodity, error) {
	list, err := u.aggregateRepo.GetAllByRange(ctx, price, size)
	if err != nil {
		return nil, err
	}

	var commodities []model.Commodity
	for _, rec := range list {
		if strings.Contains(rec.TglParsed, date) {
			commodities = append(commodities, rec)
		}
	}

	return commodities, nil
}
