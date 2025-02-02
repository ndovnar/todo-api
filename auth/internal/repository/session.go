package repository

import (
	"auth/internal/model"
	"context"
)

type SessionRepository interface {
	GetSessionByID(ctx context.Context, id string) (*model.Session, error)
	CreateSession(ctx context.Context, session *model.Session) (*model.Session, error)
	DeleteSession(ctx context.Context, id string) error
}
