package cmd

import (
	"encoding/json"
	"fmt"
	"log"

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
		Aliases: []string{"getMostRecords"},
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
