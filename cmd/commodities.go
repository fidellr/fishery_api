package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	"github.com/fidellr/fishery_api/internal/model"
	"github.com/spf13/cobra"
)

func AddRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "addCommodityRecords [data]",
		Short:   "To create more than one commodity at once",
		Aliases: []string{"createRecords"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.Commodity
			err := json.Unmarshal([]byte(data), &records)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			err = handlers.AddRecords(cmd.Context(), records)
			if err != nil {
				log.Fatalln(err)
			}

		},
	}
}

func GetAllByCommodityCmd(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getAllByCommodity",
		Short:   "To get all by given commodity name",
		Aliases: []string{"allByCommodity"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			list, err := handlers.GetAllByCommodity(cmd.Context(), data)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func GetCommodityByIDCmd(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getCommodityByID [uuid]",
		Short:   "To get commodity by given commodity ID",
		Aliases: []string{"itemById"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			list, err := handlers.GetByID(cmd.Context(), data)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}

func UpdateRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "updateCommodityRecords [payloads]",
		Short:   "To update more than one commodity at once",
		Aliases: []string{"modifyRecords"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.Commodity
			err := json.Unmarshal([]byte(data), &records)
			if err != nil {
				log.Fatalf("Unmarshaling error: %s", err.Error())
			}

			err = handlers.UpdateRecords(cmd.Context(), records)
			if err != nil {
				log.Fatalln(err)
			}

		},
	}
}

func DeleteRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "deleteCommodityRecords [uuids]",
		Short:   "To delete commodity(s) by given commodity ID more than one",
		Aliases: []string{"deleteById"},
		Run: func(cmd *cobra.Command, args []string) {
			str := args[0]

			var uuids []string
			err := json.Unmarshal([]byte(str), &uuids)
			if err != nil {
				log.Fatalln(err)
			}

			err = handlers.DeleteRecords(cmd.Context(), uuids)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
}
