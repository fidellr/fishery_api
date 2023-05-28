package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"

	"github.com/fidellr/fishery_api/internal/model"
	"github.com/fidellr/fishery_api/internal/size-options/delivery/http"
)

func AddSizeOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "addSizeOptRecords [data]",
		Short:   "To create more than one size option at once",
		Aliases: []string{"createSizeOpt"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.SizeOption
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

func UpdateSizeOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "updateSizeOptRecords [payloads]",
		Short:   "To update more than one size option at once",
		Aliases: []string{"modifySizeOptRecords"},
		Run: func(cmd *cobra.Command, args []string) {
			data := args[0]
			var records []model.SizeOption
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

func DeleteSizeOptRecords(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "deleteSizeOptRecords [payload]",
		Short:   "To delete size option(s) by given size more than one",
		Aliases: []string{"deleteSizeOpt"},
		Run: func(cmd *cobra.Command, args []string) {
			str := args[0]

			var sizeOpts []model.SizeOption
			err := json.Unmarshal([]byte(str), &sizeOpts)
			if err != nil {
				log.Fatalln(err)
			}

			err = handlers.DeleteRecords(cmd.Context(), sizeOpts)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
}
