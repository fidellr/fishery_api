package main

import (
	"context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/fidellr/fishery_api/cmd"
	commHTTP "github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	commmoditiesRepo "github.com/fidellr/fishery_api/internal/commodities/repository"
	commoditiesServices "github.com/fidellr/fishery_api/internal/commodities/usecase"

	sizeOptHTTP "github.com/fidellr/fishery_api/internal/size-options/delivery/http"
	sizeOptsRepo "github.com/fidellr/fishery_api/internal/size-options/repository"
	sizeOptsServices "github.com/fidellr/fishery_api/internal/size-options/usecase"

	areaOptHTTP "github.com/fidellr/fishery_api/internal/area-options/delivery/http"
	areaOptsRepo "github.com/fidellr/fishery_api/internal/area-options/repository"
	areaOptsServices "github.com/fidellr/fishery_api/internal/area-options/usecase"
)

func initConfig() {
	configFile := ""
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	configFile = "config.json"
	if config := viper.GetString("config"); config != "" {
		configFile = config
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalln(err.Error())
	}

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("fishery_api running in debug mode")
		return
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.Warn("fishery_api running in production mode")
}

func initCmd(ctx context.Context,
	commHandlers *commHTTP.Handler,
	sizeOptHandlers *sizeOptHTTP.Handler,
	areaOptHandlers *areaOptHTTP.Handler,
) {
	cmd.ExecuteCmd(ctx, commHandlers, sizeOptHandlers, areaOptHandlers)
}

func run() {
	initConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()

	commRepo := commmoditiesRepo.NewCommoditiesRepo()
	commSvc := commoditiesServices.NewCommoditiesUsecase(ctx, commRepo)
	aggregateRepo := commmoditiesRepo.NewAggregateRepo()
	aggregateSvc := commoditiesServices.NewAggregateUsecase(ctx, aggregateRepo)

	sizeOptRepo := sizeOptsRepo.NewSizeOptsRepo()
	sizeOptSvc := sizeOptsServices.NewCommoditiesUsecase(ctx, sizeOptRepo)

	areaOptRepo := areaOptsRepo.NewAreaOptsRepo()
	areaOptSvc := areaOptsServices.NewCommoditiesUsecase(ctx, areaOptRepo)

	initCmd(ctx,
		commHTTP.NewCommoditiesHandler(commSvc, aggregateSvc),
		sizeOptHTTP.NewSizeOptsHandler(sizeOptSvc),
		areaOptHTTP.NewAreaOptsHandler(areaOptSvc),
	)

}

func main() {
	run()
}
