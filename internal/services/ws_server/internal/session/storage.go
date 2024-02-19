package session

import (
	"context"
)

type SessionStorage interface {
	Get(ctx context.Context, id string) (Session, error)
	Add(ctx context.Context, session Session) error
	Update(ctx context.Context, session Session) error
}
