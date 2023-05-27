package cmd

import (
	"context"
	"os"

	"github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context, handlers *http.Handler) {
	rootCmd := &cobra.Command{
		Use:   "fishery",
		Short: "fishery api desc",
	}

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}

	rootCmd.AddCommand(GetAllByCommodityCmd(handlers))

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
