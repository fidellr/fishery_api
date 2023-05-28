package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	customError "github.com/pkg/errors"

	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/utils"
)

type CommoditiesRepository struct {
	client http.Client
}

type Response struct {
	*http.Response
	Body []byte
}

const (
	host              = "https://stein.efishery.com"
	searchCommodities = "/v1/storages/5e1edf521073e315924ceab4/list"
)

func NewCommoditiesRepo() *CommoditiesRepository {
	return &CommoditiesRepository{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (repo *CommoditiesRepository) AddRecords(ctx context.Context, records []model.Commodity) error {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, searchCommodities, records, http.MethodPost)
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

func (repo *CommoditiesRepository) GetAllCommodity(ctx context.Context) ([]model.Commodity, error) {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, searchCommodities, nil, http.MethodGet)
	if err != nil {
		return nil, customError.Wrap(err, "new request failing")
	}

	res, err := r.DoRequest(req)
	if err != nil {
		return nil, customError.Wrap(err, "do request failing")
	}

	if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
		return nil, err
	}

	var response []model.Commodity

	if err = json.Unmarshal(res.Body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *CommoditiesRepository) GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error) {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, searchCommodities, nil, http.MethodGet)
	if err != nil {
		return nil, customError.Wrap(err, "new request failing")
	}

	query := req.URL.Query()
	query.Add("search", fmt.Sprintf(`{"komoditas": "%s"}`, commodity))
	req.URL.RawQuery = query.Encode()

	res, err := r.DoRequest(req)
	if err != nil {
		return nil, customError.Wrap(err, "do request failing")
	}

	if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
		return nil, err
	}

	var response []model.Commodity

	if err = json.Unmarshal(res.Body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *CommoditiesRepository) GetByID(ctx context.Context, uuid string) ([]model.Commodity, error) {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, searchCommodities, nil, http.MethodGet)
	if err != nil {
		return nil, customError.Wrap(err, "new request failing")
	}

	query := req.URL.Query()
	query.Add("search", fmt.Sprintf(`{"uuid": "%s"}`, uuid))
	req.URL.RawQuery = query.Encode()

	res, err := r.DoRequest(req)
	if err != nil {
		return nil, customError.Wrap(err, "do request failing")
	}

	if err = utils.HandleHTTPError(res.StatusCode, res.Body); err != nil {
		return nil, err
	}

	fmt.Println(res.StatusCode)
	var response []model.Commodity

	if err = json.Unmarshal(res.Body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *CommoditiesRepository) UpdateRecords(ctx context.Context, payloads []model.Commodity) error {
	r := utils.Request()

	respChan := make(chan model.ResponseChannel)
	for _, payload := range payloads {
		condition := &model.UpdateCommodityQuery{
			Condition: &model.UpdateCommodityCondition{
				UUID: payload.UUID,
			},
			SetQuery: &model.SetQuery{
				Set: &model.Commodity{
					UUID:         payload.UUID,
					Komoditas:    payload.Komoditas,
					AreaProvinsi: payload.AreaProvinsi,
					AreaKota:     payload.AreaKota,
					Size:         payload.Size,
					Price:        payload.Price,
					TglParsed:    payload.TglParsed,
					Timestamp:    payload.Timestamp,
				},
			},
		}

		go func() {
			req, err := r.NewRequest(ctx, host, searchCommodities, condition, http.MethodPut)
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

func (repo *CommoditiesRepository) DeleteRecords(ctx context.Context, uuid []string) error {
	r := utils.Request()

	respChan := make(chan model.ResponseChannel)
	for _, id := range uuid {
		condition := &model.DeleteCommodityQuery{
			Condition: &model.DeleteCommodityCondition{
				UUID: id,
			},
		}

		go func() {
			req, err := r.NewRequest(ctx, host, searchCommodities, condition, http.MethodDelete)
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
