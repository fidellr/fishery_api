package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/spf13/cobra"
)

func GetMostCommodityRecordsCmd(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getMostCommodityRecords",
		Short:   "To get the most commodity records",
		Aliases: []string{"getMostRecords"},
		Run: func(cmd *cobra.Command, args []string) {
			list, err := handlers.GetMostCommodityRecords(cmd.Context())
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func GetByArea(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getCommodityByArea [data]",
		Short:   "To get commodity records by the given area",
		Aliases: []string{"getByArea"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]

			var record model.Commodity
			err := json.Unmarshal([]byte(data), &record)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			list, err := handlers.GetByArea(cmd.Context(), record.AreaKota, record.AreaProvinsi)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func GetMaxPrice(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getMaxPrice [week, commodity]",
		Short:   "To get the max price of commodity records by the given week and commodity",
		Aliases: []string{"getMax"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]

			var record map[string]string
			err := json.Unmarshal([]byte(data), &record)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			week, err := strconv.ParseInt(record["week"], 10, 64)
			if err != nil {
				log.Fatalln(err)
			}

			list, err := handlers.GetMaxPrice(cmd.Context(), int(week), record["komoditas"])
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func GetAllByRange(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getAllByRange [price, size, date=YYYY-MM-DD]",
		Short:   "To get commodity by the given price, size, and date",
		Aliases: []string{"getByRange"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]

			var record map[string]string
			err := json.Unmarshal([]byte(data), &record)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			list, err := handlers.GetAllByRange(cmd.Context(), record["price"], record["size"], record["date"])
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func GetAllByCommodityAndArea(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getAllByCommodityAndArea [komoditas, area_provinsi, area_kota]",
		Short:   "To get commodity by the given commodity name and area",
		Aliases: []string{"getByCommodityAndArea"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]

			var dataJson map[string]string
			err := json.Unmarshal([]byte(data), &dataJson)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			listStr, err := handlers.GetAllByCommodity(cmd.Context(), dataJson["komoditas"])
			if err != nil {
				log.Fatalln(err)
			}

			var records []model.Commodity
			err = json.Unmarshal([]byte(listStr), &records)
			if err != nil {
				log.Fatalln(err)
			}

			for i := len(records) - 1; i >= 0; i-- {
				rec := &records[i]

				if rec.AreaKota != dataJson["area_kota"] && rec.AreaProvinsi != dataJson["area_provinsi"] {
					records = append(records[:i], records[i+1:]...)
				}
			}

			jsonByte, err := json.MarshalIndent(records, "", "    ")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(string(jsonByte))
		},
	}
}

func GetAllByPriceRange(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getAllByPriceRange [startPrice-endPrice]",
		Short:   "To get commodity by the given price range",
		Aliases: []string{"getByRange"},
		Run: func(cmd *cobra.Command, args []string) {
			dataJson := args[0]

			var record map[string]string
			err := json.Unmarshal([]byte(dataJson), &record)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			listStr, err := handlers.GetAllCommodity(cmd.Context())
			if err != nil {
				log.Fatalln(err)
			}

			var records []model.Commodity
			err = json.Unmarshal([]byte(listStr), &records)
			if err != nil {
				log.Fatalln(err)
			}

			if !strings.Contains(record["price"], "-") {
				log.Fatalln("Invalid input, price should in startPrice-endPrice format")
			}

			partsPrice := strings.Split(record["price"], "-")

			if partsPrice[0] == "" || partsPrice[1] == "" || len(partsPrice) != 2 {
				log.Fatalln("Invalid input, price should contain startPrice-endPrice")
			}

			startPrice, err := strconv.ParseInt(partsPrice[0], 10, 64)
			if err != nil {
				log.Fatalln(err)
			}
			endPrice, err := strconv.ParseInt(partsPrice[1], 10, 64)
			if err != nil {
				log.Fatalln(err)
			}

			var matchedRecords []model.Commodity
			for _, rec := range records {
				if len(rec.Price) > 0 {
					recPrice, err := strconv.ParseInt(rec.Price, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}

					if recPrice >= startPrice && recPrice <= endPrice {
						matchedRecords = append(matchedRecords, rec)
					}
				}
			}

			jsonByte, err := json.MarshalIndent(matchedRecords, "", "    ")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(string(jsonByte))
		},
	}
}

func GetLatestTenCommodities(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getLatestTenCommodities",
		Short:   "To get commodity by the given price, size, and date",
		Aliases: []string{"getLatestTenCommodities"},
		Run: func(cmd *cobra.Command, args []string) {
			listStr, err := handlers.GetAllCommodity(cmd.Context())
			if err != nil {
				log.Fatalln(err)
			}

			var records []model.Commodity
			err = json.Unmarshal([]byte(listStr), &records)
			if err != nil {
				log.Fatalln(err)
			}

			length := len(records)
			startIndex := length - 10
			if startIndex < 0 {
				startIndex = 0
			}
			last10Records := records[startIndex:]

			jsonByte, err := json.MarshalIndent(last10Records, "", "    ")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(string(jsonByte))
		},
	}
}
