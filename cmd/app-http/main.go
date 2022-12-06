package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/server"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/repo/postgresql"
	"github.com/resyahrial/go-template/pkg/graceful"
	"github.com/resyahrial/go-template/pkg/logger"
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

	if !config.GlobalConfig.App.DebugMode {
		logger.UseZapLogger(
			logger.WithCore(
				logger.ZapLoggerReleaseModeCore(),
			),
		)
	}

	_ = postgresql.InitDatabase(config.GlobalConfig)

	svr := &http.Server{
		Addr: fmt.Sprintf(":%v", config.GlobalConfig.App.ServerAppPort),
	}

	serverOpts := []server.Option{
		handler.New(
			handler.WithGorm(postgresql.DbInstance),
		),
	}

	if config.GlobalConfig.App.DebugMode {
		svr.Handler = server.InitServerDebugMode(serverOpts...)
	} else {
		svr.Handler = server.InitServerReleaseMode(serverOpts...)
	}

	graceful.RunHttpServer(context.Background(), svr, 10*time.Second)
}
