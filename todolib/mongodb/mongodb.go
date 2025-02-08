package mongodb

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const clientTimeoutShutdown = 20 * time.Second

type DB struct {
	DB *mongo.Database
}

func New(ctx context.Context, config *Config, appName string) (*DB, error) {
	connectionDebugInfo := databaseConnectionInfoString(config)

	clientOptions := options.Client().
		SetHosts(config.Hosts).
		SetAuth(options.Credential{
			AuthSource: config.Database,
			Username:   config.Username,
			Password:   config.Password,
		}).
		SetConnectTimeout(5 * time.Second).
		SetAppName(appName)

	if config.UseTLS {
		clientOptions.SetTLSConfig(&tls.Config{MinVersion: tls.VersionTLS12})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		err := fmt.Errorf("could not create MongoDB client [%v]: %w", connectionDebugInfo, err)
		return nil, err
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	err = client.Ping(ctxTimeout, nil)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("connection to MongoDB [%v] worked, but ping failed: %w", connectionDebugInfo, err)
	}

	return &DB{
		DB: client.Database(config.Database),
	}, nil
}

func (d *DB) RunShutdown(ctx context.Context) error {
	<-ctx.Done()

	ctxTimeout, cancel := context.WithTimeout(context.Background(), clientTimeoutShutdown)
	defer cancel()
	if err := d.DB.Client().Disconnect(ctxTimeout); err != nil {
		return fmt.Errorf("disconnecting: %w", err)
	}

	return nil
}

func databaseConnectionInfoString(config *Config) string {
	return fmt.Sprintf(
		"mongodb://%s@%s/%s",
		config.Username,
		strings.Join(config.Hosts, ","),
		config.Database,
	)
}
