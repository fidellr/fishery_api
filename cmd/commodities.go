package cmd

import (
	"fmt"
	"log"

	"github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	"github.com/spf13/cobra"
)

func GetAllByCommodityCmd(handlers *http.Handler) *cobra.Command {
	return &cobra.Command{
		Use:     "getAllByCommodity",
		Short:   "To get all by given commodity name",
		Aliases: []string{"allByCommodity"},
		Run: func(cmd *cobra.Command, args []string) {
			list, err := handlers.GetAllByCommodity(cmd.Context(), args[0])
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(list)
		},
	}
}
