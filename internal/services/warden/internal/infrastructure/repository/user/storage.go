package auth

import (
	"context"
	"database/sql"

	"github.com/optclblast/biocom/internal/services/warden/internal/lib/models"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/sqlutils"
	"github.com/optclblast/biocom/internal/services/warden/internal/usecase/repository/user"
)

type UserSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *UserSQL {
	return &UserSQL{db}
}

type transactionCtxKey struct{}

func (s *UserSQL) Conn(ctx context.Context) sqlutils.DBTX {
	if tx, ok := ctx.Value(transactionCtxKey{}).(*sql.Tx); ok {
		return tx
	}

	return s.db
}

func (s *UserSQL) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return nil // todo
}

func (r *UserSQL) Create(ctx context.Context, params user.CreateParams) error { return nil }

func (r *UserSQL) Update(ctx context.Context, params user.UpdateParams) error { return nil }

func (r *UserSQL) Delete(ctx context.Context, params user.DeleteParams) error { return nil }

func (r *UserSQL) Get(ctx context.Context, params user.GetParams) (*models.User, error) {
	return nil, nil
}

func (r *UserSQL) GetUsers(ctx context.Context, params user.GetUsersParams) ([]*models.User, error) {
	return nil, nil
}
