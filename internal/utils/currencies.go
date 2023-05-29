package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	url    = "https://api.freecurrencyapi.com/v1/latest"
	apiKey = "apikey=xyUu7v1ZIwqG0yNDxoXZdaFzKRHmxiQSIOkRhYiK"
)

type CurrencyConverter struct {
	cache    map[string]float64
	cacheMux sync.RWMutex
}

func NewCurrencyConverter() *CurrencyConverter {
	return &CurrencyConverter{
		cache: make(map[string]float64),
	}
}

func (cc *CurrencyConverter) ConvertToUSD(price float64, currency string) (float64, error) {
	cc.cacheMux.RLock()
	rate, found := cc.cache[currency]
	cc.cacheMux.RUnlock()

	if !found {
		rate, err := cc.fetchConversionRateToUSD(currency)
		if err != nil {
			return 0, err
		}

		cc.cacheMux.Lock()
		cc.cache[currency] = rate
		cc.cacheMux.Unlock()
	}

	convertedPrice := price / rate
	return convertedPrice, nil
}

type ConversionResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func (cc *CurrencyConverter) fetchConversionRateToUSD(currency string) (float64, error) {
	response, err := http.Get(fmt.Sprintf("%s?%s&currencies=USD&base_currency=%s", url, apiKey, currency))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var data struct {
		Data struct {
			Usd float64 `json:"USD"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	rate := data.Data.Usd
	if rate == 0 {
		return 0, fmt.Errorf("conversion rate not found for currency %s", currency)
	}

	return rate, nil
}
