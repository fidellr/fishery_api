package repository

import (
	"context"
	"fmt"
	"net/http"
	"time"

	customError "github.com/pkg/errors"

	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type AreaOptionsRepository struct {
	client http.Client
}

type Response struct {
	*http.Response
	Body []byte
}

const (
	host = "https://stein.efishery.com"
	path = "/v1/storages/5e1edf521073e315924ceab4/option_area"
)

func NewAreaOptsRepo() *AreaOptionsRepository {
	return &AreaOptionsRepository{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (repo *AreaOptionsRepository) AddRecords(ctx context.Context, records []model.AreaOption) error {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, path, records, http.MethodPost)
	if err != nil {
		return customError.Wrap(err, "new request failing")
	}

	res, err := r.DoRequest(req)
	if err != nil {
		return customError.Wrap(err, "do request failing")
	}

	if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
		return err
	}

	return nil
}

func (repo *AreaOptionsRepository) UpdateRecords(ctx context.Context, payloads []model.AreaOption) error {
	r := utils.Request()

	respChan := make(chan model.ResponseChannel)
	for _, payload := range payloads {
		condition := &model.UpdateAreaOptQuery{
			Condition: &model.UpdateAreaOptCondition{
				Province: payload.Province,
			},
			SetQuery: &model.SetQuery{
				Set: &model.AreaOption{
					Province: payload.Province,
					City:     payload.City,
				},
			},
		}

		go func() {
			req, err := r.NewRequest(ctx, host, path, condition, http.MethodPut)
			if err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   customError.Wrap(err, "new request failing"),
				}
			}

			res, err := r.DoRequest(req)
			if err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   customError.Wrap(err, "do request failing"),
				}
			}
			fmt.Println(res.StatusCode)

			if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   err,
				}
			}

			respChan <- model.ResponseChannel{
				Value: res,
				Err:   nil,
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			return utils.ConstraintErrorf("fetch tooks too long")
		case resp := <-respChan:
			if resp.Err != nil {
				return resp.Err
			}

			return nil
		}
	}
}

func (repo *AreaOptionsRepository) DeleteRecords(ctx context.Context, records []model.AreaOption) error {
	r := utils.Request()

	respChan := make(chan model.ResponseChannel)
	for _, opt := range records {
		condition := &model.DeleteAreaOptQuery{
			Condition: &model.DeleteAreaOptCondition{
				Province: opt.Province,
				City:     opt.City,
			},
		}

		go func() {
			req, err := r.NewRequest(ctx, host, path, condition, http.MethodDelete)
			if err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   customError.Wrap(err, "new request failing"),
				}
			}

			res, err := r.DoRequest(req)
			if err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   customError.Wrap(err, "do request failing"),
				}
			}
			fmt.Println(res.StatusCode)

			if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
				respChan <- model.ResponseChannel{
					Value: nil,
					Err:   err,
				}
			}

			respChan <- model.ResponseChannel{
				Value: res,
				Err:   nil,
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			return utils.ConstraintErrorf("fetch tooks too long")
		case resp := <-respChan:
			if resp.Err != nil {
				return resp.Err
			}

			return nil
		}
	}
}
