package usecase

import (
	"context"
	"fmt"

	areaoptions "github.com/fidellr/fishery_api/internal/area-options"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type areaOptUsecase struct {
	areaOptRepo areaoptions.AreaOptionRepo
	context     context.Context
}

func NewCommoditiesUsecase(ctx context.Context, areaOptRepo areaoptions.AreaOptionRepo) areaoptions.AreaOptionUsecase {
	return &areaOptUsecase{
		areaOptRepo: areaOptRepo,
		context:     ctx,
	}
}

func (u *areaOptUsecase) AddRecords(ctx context.Context, records []model.AreaOption) error {
	if records == nil || len(records) < 1 {
		return utils.ConstraintErrorf("Data cannot be empty or nil")
	}

	for i := range records {
		rec := &records[i]
		if err := utils.Validate(rec); err != nil {
			return err
		}
	}
	fmt.Println(records)

	err := u.areaOptRepo.AddRecords(ctx, records)
	if err != nil {
		return err
	}

	return nil
}

func (u *areaOptUsecase) UpdateRecords(ctx context.Context, payloads []model.AreaOption) error {
	if payloads == nil || len(payloads) < 1 {
		return utils.ConstraintErrorf("Data cannot be empty or nil")
	}

	for i := range payloads {
		payload := &payloads[i]
		if err := utils.Validate(payload); err != nil {
			return err
		}
	}
	fmt.Println(payloads)

	err := u.areaOptRepo.UpdateRecords(ctx, payloads)
	if err != nil {
		return err
	}

	return nil
}

func (u *areaOptUsecase) DeleteRecords(ctx context.Context, records []model.AreaOption) error {
	if len(records) < 1 {
		return utils.ConstraintErrorf("UUID cannot be empty")
	}

	err := u.areaOptRepo.DeleteRecords(ctx, records)
	if err != nil {
		return err
	}

	return nil
}
