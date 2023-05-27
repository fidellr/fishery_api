package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/fidellr/fishery_api/cmd"
	commHTTP "github.com/fidellr/fishery_api/internal/commodities/delivery/http"
	commmoditiesRepo "github.com/fidellr/fishery_api/internal/commodities/repository"
	commoditiesServices "github.com/fidellr/fishery_api/internal/commodities/usecase"
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

func initCmd(ctx context.Context, handlers *commHTTP.Handler) {
	cmd.Execute(ctx, handlers)
}

func run() {
	initConfig()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()

	commRepo := commmoditiesRepo.NewCommoditiesRepo()
	commSvc := commoditiesServices.NewCommoditiesUsecase(ctx, commRepo)
	initCmd(ctx, commHTTP.NewCommoditiesHandler(commSvc))
}

func main() {

	run()
	fmt.Println("TEST")
}
