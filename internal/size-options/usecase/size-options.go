package usecase

import (
	"context"
	"fmt"

	"github.com/fidellr/fishery_api/internal/model"
	sizeoptions "github.com/fidellr/fishery_api/internal/size-options"
	"github.com/fidellr/fishery_api/internal/utils"
)

type sizeOptUsecase struct {
	sizeOptRepo sizeoptions.SizeOptionRepo
	context     context.Context
}

func NewCommoditiesUsecase(ctx context.Context, sizeOptRepo sizeoptions.SizeOptionRepo) sizeoptions.SizeOptionUsecase {
	return &sizeOptUsecase{
		sizeOptRepo: sizeOptRepo,
		context:     ctx,
	}
}

func (u *sizeOptUsecase) AddRecords(ctx context.Context, records []model.SizeOption) error {
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

	err := u.sizeOptRepo.AddRecords(ctx, records)
	if err != nil {
		return err
	}

	return nil
}

func (u *sizeOptUsecase) UpdateRecords(ctx context.Context, payloads []model.SizeOption) error {
	if payloads == nil || len(payloads) < 1 {
		return utils.ConstraintErrorf("Data cannot be empty or nil")
	}

	for i := range payloads {
		payload := &payloads[i]
		if err := utils.Validate(payload); err != nil {
			return err
		}
	}

	err := u.sizeOptRepo.UpdateRecords(ctx, payloads)
	if err != nil {
		return err
	}

	return nil
}

func (u *sizeOptUsecase) DeleteRecords(ctx context.Context, records []model.SizeOption) error {
	if len(records) < 1 {
		return utils.ConstraintErrorf("UUID cannot be empty")
	}

	err := u.sizeOptRepo.DeleteRecords(ctx, records)
	if err != nil {
		return err
	}

	return nil
}
