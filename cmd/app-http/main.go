package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/resyahrial/go-template/config"
	route "github.com/resyahrial/go-template/internal/api/routes"
	"github.com/resyahrial/go-template/internal/api/server"
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

	serverEngine := server.InitGinEngine(config.GlobalConfig.App)
	if serverEngine == nil {
		log.Fatal("server failed to initialized")
	}

	go func() {
		// run http connections
		log.Printf("Running http server on port : %v", config.GlobalConfig.App.ServerAppPort)
		graceful.RunHttpServer(context.Background(), &http.Server{
			Addr:    fmt.Sprintf(":%v", config.GlobalConfig.App.ServerAppPort),
			Handler: route.InitRoutes(serverEngine),
		}, 10*time.Second)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		<-sigs
		// worker.Shutdown()
		done <- true
	}()
	<-done

}
