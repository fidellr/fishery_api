package cmd

import (
	"context"
	"os"

	areaOptHttp "github.com/fidellr/fishery_api/internal/area-options/delivery/http"
	commodityHttp "github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	sizeOptHttp "github.com/fidellr/fishery_api/internal/size-options/delivery/http"
	"github.com/spf13/cobra"
)

func ExecuteCmd(
	ctx context.Context,
	commHandlers *commodityHttp.Handler,
	sizeOptHandlers *sizeOptHttp.Handler,
	areaOptHandlers *areaOptHttp.Handler,
) {
	cmd := &cobra.Command{
		Use:   "fishery",
		Short: "fishery api to create, get or modify commodity and options",
	}

	err := cmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}

	cmd.AddCommand(
		AddRecords(commHandlers),
		GetAllByCommodityCmd(commHandlers),
		GetCommodityByIDCmd(commHandlers),
		DeleteRecords(commHandlers),
		UpdateRecords(commHandlers),

		GetMostCommodityRecordsCmd(commHandlers),
		GetByArea(commHandlers),

		AddSizeOptRecords(sizeOptHandlers),
		UpdateSizeOptRecords(sizeOptHandlers),
		DeleteSizeOptRecords(sizeOptHandlers),

		AddAreaOptRecords(areaOptHandlers),
		UpdateAreaOptRecords(areaOptHandlers),
		DeleteAreaOptRecords(areaOptHandlers))

	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
