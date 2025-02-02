package main

import (
	"context"
	"errors"
	"lib/mongodb"
	"lib/pem"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/httpapi"
	"auth/internal/repository/mongo"
	"auth/internal/service"
	"auth/internal/token"
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

	privateKey, err := pem.ReadPriviateKey(config.PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read private key")
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

	err = db.CreateIndexes(sigCtx, mongodb.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create indexes")
	}

	tokenMaker := token.NewTokenMaker(&config.Token, privateKey, publicKey)

	userRepository := mongo.NewUserRepository(mongodb.DB)
	sessionsRepository := mongo.NewSessionRepository(mongodb.DB)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(tokenMaker, sessionsRepository, userRepository)

	httpapi := httpapi.New(&config.HTTPAPI, publicKey, userService, authService)

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
