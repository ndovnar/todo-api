package repository

import (
	"context"

	"todoauth/internal/modeldb"
)

type SessionRepository interface {
	GetSessionByID(ctx context.Context, id string) (*modeldb.Session, error)
	CreateSession(ctx context.Context, userID string) (*modeldb.Session, error)
	DeleteSession(ctx context.Context, id string) error
}
