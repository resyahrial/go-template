package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/server"
	"github.com/resyahrial/go-template/internal/repositories/postgresql"
	"github.com/resyahrial/go-template/pkg/graceful"
)

type (
	Flag struct {
		Environment string
	}
)

var (
	appFlag Flag
)

func init() {
	flag.StringVar(
		&appFlag.Environment,
		"env",
		"dev",
		"env of deployment, will load the respective yml conf file.",
	)
}

func main() {
	flag.Parse()
	config.InitConfig(appFlag.Environment)

	_ = postgresql.InitDatabase(config.GlobalConfig)

	graceful.RunHttpServer(context.Background(), &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GlobalConfig.App.ServerAppPort),
		Handler: server.InitServer(postgresql.DbInstance, config.GlobalConfig),
	}, 10*time.Second)
}
