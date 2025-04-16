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

	"todostatistic/internal/config"
	"todostatistic/internal/consumer"
	"todostatistic/internal/httpapi"
	"todostatistic/internal/repository/mongo"
	"todostatistic/internal/service"
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

	publicKey, err := pem.DecodePublicKey(config.PublicKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to decode public key")
	}

	group, errCtx := errgroup.WithContext(sigCtx)

	mongodb, err := mongodb.New(sigCtx, &config.Mongo, application)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create db")
	}

	group.Go(func() error {
		return mongodb.RunShutdown(errCtx)
	})

	statisticRepository := mongo.NewStatisticRepository(mongodb.DB)
	statisticService := service.NewStatisticService(statisticRepository)

	consumer, err := consumer.NewConsumer(&config.Consumer, statisticService)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create consumer")
	}
	group.Go(consumer.Run)
	group.Go(func() error {
		return consumer.RunShutdown(errCtx)
	})

	httpapi := httpapi.New(&config.HTTPAPI, publicKey, statisticService)
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
