package cmd

/**
TODO:
1. getAllByRange (harga, size, tanggal)
2. getMaxPrice (by week, by commodity)
3. Mencari berdasar range harga
4. 10 Latest Data (switch index aja dan take 10 items)
5. Tambahkan informasi harga dalam USD dengan memanfaatkan layanan
currency converter
	1. bikin modulenya sendiri dalam aplikasi dan terapkan mekanisme caching
**/
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
		Short: "fishery cli to create, get, aggregate or modify commodity and options",
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
		GetMaxPrice(commHandlers),
		GetAllByRange(commHandlers),
		GetAllByCommodityAndArea(commHandlers),
		GetAllByPriceRange(commHandlers),
		GetLatestTenCommodities(commHandlers),

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
