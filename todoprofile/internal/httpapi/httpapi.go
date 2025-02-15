package httpapi

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"todoprofile/internal/service"
)

const shutdownMaxDuration = 15 * time.Second

type HTTPAPI struct {
	config         *Config
	publicKey      *rsa.PublicKey
	server         *http.Server
	router         *gin.Engine
	profileService service.ProfileService
}

func New(
	config *Config,
	publicKey *rsa.PublicKey,
	profileService service.ProfileService,
) *HTTPAPI {
	router := gin.Default()

	api := &HTTPAPI{
		config:         config,
		publicKey:      publicKey,
		profileService: profileService,
		router:         router,
		server: &http.Server{
			Addr:              fmt.Sprintf(":%v", config.Port),
			ReadHeaderTimeout: 5 * time.Second,
			Handler:           router.Handler(),
		},
	}

	api.regiesterRoutes()

	return api
}

func (a *HTTPAPI) Run() error {
	log.Info().Msgf("api - listening on port %v", a.config.Port)

	if err := a.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("listening and serving: %w", err)
	}

	return nil
}

func (a *HTTPAPI) RunShutdown(ctx context.Context) error {
	<-ctx.Done()

	ctxStop, cancel := context.WithTimeout(context.Background(), shutdownMaxDuration)
	defer cancel()

	if err := a.server.Shutdown(ctxStop); err != nil {
		return err
	}

	return nil
}
