package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fidellr/fishery_api/internal/commodities"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
	"github.com/google/uuid"
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

func (u *commoditiesUsecase) AddRecords(ctx context.Context, records []model.Commodity) error {
	cc := utils.NewCurrencyConverter()
	if records == nil || len(records) < 1 {
		return utils.ConstraintErrorf("Data cannot be empty or nil")
	}

	for i := range records {
		rec := &records[i]
		rec.UUID = uuid.New().String()
		rec.Komoditas = strings.ToUpper(rec.Komoditas)
		rec.AreaProvinsi = strings.ToUpper(rec.AreaProvinsi)
		rec.AreaKota = strings.ToUpper(rec.AreaKota)
		rec.TglParsed = utils.ParseTime(time.Now())
		rec.Timestamp = fmt.Sprintf("%v", utils.TimeIn(time.Now(), "Indonesia").UnixMilli())

		recPrice, err := strconv.ParseFloat(rec.Price, 64)
		if err != nil {
			return err
		}
		convertedPrice, err := cc.ConvertToUSD(recPrice, "IDR")
		if err != nil {
			return err
		}

		rec.USD = fmt.Sprintf("%v", convertedPrice)
		if err := utils.Validate(rec); err != nil {
			return err
		}
	}
	fmt.Println(records)

	err := u.commoditiesRepo.AddRecords(ctx, records)
	if err != nil {
		return err
	}

	return nil
}

func convertPriceToUSD(cc *utils.CurrencyConverter, price float64, currency string) (float64, error) {
	convertedPrice, err := cc.ConvertToUSD(price, currency)
	if err != nil {
		return 0, err
	}

	return convertedPrice, nil
}

func (u *commoditiesUsecase) GetAllCommodity(ctx context.Context) ([]model.Commodity, error) {
	list, err := u.commoditiesRepo.GetAllCommodity(ctx)
	if err != nil {
		return nil, err
	}

	cc := utils.NewCurrencyConverter()
	for i := range list {
		rec := &list[i]
		if rec.Price != "" {
			recPrice, err := strconv.ParseFloat(rec.Price, 64)
			if err != nil {
				return nil, err
			}

			convertedPrice, err := convertPriceToUSD(cc, recPrice, "IDR")
			if err != nil {
				return nil, err
			}

			rec.USD = fmt.Sprintf("%v", convertedPrice)
			if err := utils.Validate(rec); err != nil {
				return nil, err
			}
		}
	}

	return list, nil
}

func (u *commoditiesUsecase) GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error) {
	list, err := u.commoditiesRepo.GetAllByCommodity(ctx, strings.ToUpper(commodity))
	if err != nil {
		return nil, err
	}

	cc := utils.NewCurrencyConverter()
	for i := range list {
		rec := &list[i]
		recPrice, err := strconv.ParseFloat(rec.Price, 64)
		if err != nil {
			return nil, err
		}

		convertedPrice, err := convertPriceToUSD(cc, recPrice, "IDR")
		if err != nil {
			return nil, err
		}

		rec.USD = fmt.Sprintf("%v", convertedPrice)
		if err := utils.Validate(rec); err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (u *commoditiesUsecase) GetByID(ctx context.Context, uuid string) ([]model.Commodity, error) {
	list, err := u.commoditiesRepo.GetByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	cc := utils.NewCurrencyConverter()
	for i := range list {
		rec := &list[i]
		recPrice, err := strconv.ParseFloat(rec.Price, 64)
		if err != nil {
			return nil, err
		}

		convertedPrice, err := convertPriceToUSD(cc, recPrice, "IDR")
		if err != nil {
			return nil, err
		}

		rec.USD = fmt.Sprintf("%v", convertedPrice)
		if err := utils.Validate(rec); err != nil {
			return nil, err
		}
	}

	return list, err
}

func (u *commoditiesUsecase) UpdateRecords(ctx context.Context, payloads []model.Commodity) error {
	if payloads == nil || len(payloads) < 1 {
		return utils.ConstraintErrorf("Data cannot be empty or nil")
	}

	for i := range payloads {
		payload := &payloads[i]
		payload.Komoditas = strings.ToUpper(payload.Komoditas)
		payload.AreaProvinsi = strings.ToUpper(payload.AreaProvinsi)
		payload.AreaKota = strings.ToUpper(payload.AreaKota)
		payload.Timestamp = fmt.Sprintf("%v", utils.TimeIn(time.Now(), "Indonesia").UnixMilli())

		if err := utils.Validate(payload); err != nil {
			return err
		}
	}
	fmt.Println(payloads)

	err := u.commoditiesRepo.UpdateRecords(ctx, payloads)
	if err != nil {
		return err
	}

	return nil
}

func (u *commoditiesUsecase) DeleteRecords(ctx context.Context, uuid []string) error {
	if len(uuid) < 1 {
		return utils.ConstraintErrorf("UUID cannot be empty")
	}

	err := u.commoditiesRepo.DeleteRecords(ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}
