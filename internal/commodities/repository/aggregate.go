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

type AggregateRepository struct {
	client http.Client
}

func NewAggregateRepo() *AggregateRepository {
	return &AggregateRepository{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func request(ctx context.Context) ([]model.Commodity, error) {
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

func (repo *AggregateRepository) GetMostCommodityRecords(ctx context.Context) ([]model.Commodity, error) {
	commRepo := NewCommoditiesRepo()
	list, err := commRepo.GetAllCommodity(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) < 1 {
		return nil, utils.ConstraintErrorf("Your requested item is not exist")
	}

	return request(ctx)
}

func (repo *AggregateRepository) GetByArea(ctx context.Context, city, province string) ([]model.Commodity, error) {
	r := utils.Request()

	req, err := r.NewRequest(ctx, host, searchCommodities, nil, http.MethodGet)
	if err != nil {
		return nil, customError.Wrap(err, "new request failing")
	}

	query := req.URL.Query()
	query.Add("search", fmt.Sprintf(`{"area_provinsi": "%s", "area_kota": "%s"}`, province, city))
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
