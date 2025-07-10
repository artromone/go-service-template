// Package app configures and runs application.
package app

import (
	"fmt"
	"go-clean-template/config"
	"go-clean-template/internal/controller/grpc"
	"go-clean-template/internal/controller/http"
	"go-clean-template/internal/repo/persistent"
	"go-clean-template/internal/repo/webapi"
	"go-clean-template/internal/usecase/translation"
	"go-clean-template/pkg/grpcserver"
	"go-clean-template/pkg/httpserver"
	"go-clean-template/pkg/logger"
	"go-clean-template/pkg/postgres"
	"go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	"os"
	"os/signal"
	"syscall"

	amqprpc "go-clean-template/internal/controller/amqp_rpc"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use-Case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase, log)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, log)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}
	rmqServer.Start()

	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.App, translationUseCase, log)
	grpcServer.Start()

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, cfg, translationUseCase, log)
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		log.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		log.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = grpcServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
