package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"todolib/mongodb"
	"todolib/pem"

	"todo/internal/config"
	"todo/internal/httpapi"
	"todo/internal/repository/mongo"
	"todo/internal/service"
)

var application string

func main() {
	sigCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCtx.Done()
		log.Info().Msg("shutdown signal received - attempting graceful shutdown")
		cancel()
	}()

	config, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	publicKey, err := pem.ReadPublicKey(config.PublicKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read public key")
	}

	group, errCtx := errgroup.WithContext(sigCtx)

	mongodb, err := mongodb.New(sigCtx, &config.Mongo, application)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create db")
	}

	group.Go(func() error {
		return mongodb.RunShutdown(errCtx)
	})

	todoRepository := mongo.NewTodoRepository(mongodb.DB)
	todoService := service.NewTodoService(todoRepository)

	httpapi := httpapi.New(&config.HTTPAPI, publicKey, todoService)

	group.Go(httpapi.Run)
	group.Go(func() error {
		return httpapi.RunShutdown(errCtx)
	})

	if err := group.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Fatal().Err(err).Msg("main - shutdown completed with error(s)")
		}
	}

	log.Info().Msg("main - shutdown completed without errors")
}
