package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"

	"github.com/fidellr/fishery_api/internal/area-options/delivery/http"
	"github.com/fidellr/fishery_api/internal/model"
)

func AddAreaOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "addAreaOptRecords [data]",
		Short:   "To create more than one area option at once",
		Aliases: []string{"createAreaOpt"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.AreaOption
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

func UpdateAreaOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "updateAreaOptRecords [payloads]",
		Short:   "To update more than one area option at once",
		Aliases: []string{"modifyAreaOptRecords"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.AreaOption
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

func DeleteAreaOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "deleteAreaOptRecords [payload]",
		Short:   "To delete area option(s) by given area option more than one",
		Aliases: []string{"deleteAreaOpt"},
		Run: func(cmd *cobra.Command, args []string) {
			str := args[0]

			var AreaOpts []model.AreaOption
			err := json.Unmarshal([]byte(str), &AreaOpts)
			if err != nil {
				log.Fatalln(err)
			}

			err = handlers.DeleteRecords(cmd.Context(), AreaOpts)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
}
