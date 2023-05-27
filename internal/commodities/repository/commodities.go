package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fidellr/fishery_api/internal/model"
)

type CommoditiesRepository struct {
}

type Response struct {
	res []model.Commodity
	err error
}

func NewCommoditiesRepo() *CommoditiesRepository {
	return &CommoditiesRepository{}
}

func (r *CommoditiesRepository) GetAllByCommodity(ctx context.Context, commodity string) ([]model.Commodity, error) {
	respChan := make(chan Response)
	go func() {
		res, err := http.Get(fmt.Sprintf(`https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list?search={"komoditas": "%s"}`, strings.ToUpper(commodity)))
		if err != nil {
			respChan <- Response{
				res: nil,
				err: err,
			}
		}
		defer res.Body.Close()

		var response []model.Commodity

		body, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &response)

		fmt.Printf("testttt %v", response)

		respChan <- Response{
			res: response,
			err: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("Fetch to stein takes too long")
		case resp := <-respChan:
			return resp.res, resp.err
		}
	}
}
