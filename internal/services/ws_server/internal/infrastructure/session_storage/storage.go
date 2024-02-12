package sessionstorage

import (
	"context"

	"github.com/optclblast/biocom/internal/services/ws_server/internal/session"
)

type SessionStorage interface {
	Get(ctx context.Context, id string) (session.Session, error)
	Add(ctx context.Context, session session.Session) error
	Update(ctx context.Context, session session.Session) error
}
