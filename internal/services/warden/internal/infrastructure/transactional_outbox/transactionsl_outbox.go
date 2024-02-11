package transactionaloutbox

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/events"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/sqlutils"
)

type TransactionalOutbox interface {
	Append(ctx context.Context, events ...events.Event) error
}

func NewTransactionalOutbox(
	db *sql.DB,
	log *slog.Logger,
	eventsInteractor events.EventsInteractor,
) TransactionalOutbox {
	txOutox := &transactionalOutboxPostgres{
		chron: &transactionalOutboxCron{
			log:              log,
			eventsInteractor: eventsInteractor,
			fn: func() {

			},
		},
		db: db,
	}

	return txOutox
}

type transactionalOutboxPostgres struct {
	db    *sql.DB
	chron *transactionalOutboxCron
}

type transactionCtxKey struct{}

func (s *transactionalOutboxPostgres) conn(ctx context.Context) sqlutils.DBTX {
	if tx, ok := ctx.Value(transactionCtxKey{}).(*sql.Tx); ok {
		return tx
	}

	return s.db
}

func (s *transactionalOutboxPostgres) transaction(
	ctx context.Context,
	f func(ctx context.Context) error,
) error {
	var err error
	var tx *sql.Tx = new(sql.Tx)

	defer func() {
		if err == nil {
			err = tx.Commit()
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = errors.Join(fmt.Errorf("error rollback transaction: %w", rbErr), err)
				return
			}

			err = fmt.Errorf("error commit transaction: %w", err)
		}
	}()

	if _, ok := ctx.Value(transactionCtxKey{}).(*sql.Tx); !ok {
		tx, err = s.db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelRepeatableRead,
		})
		if err != nil {
			return fmt.Errorf("error begin transaction: %w", err)
		}

		ctx = context.WithValue(ctx, transactionCtxKey{}, tx)
	}

	err = f(ctx)
	if err != nil {
		return fmt.Errorf("error run transaction function: %w", err)
	}

	return err
}

func (o *transactionalOutboxPostgres) Append(ctx context.Context, events ...events.Event) error {
	err := o.transaction(ctx, func(ctx context.Context) error {
		for _, e := range events {
			query := sq.Insert("transactional_outbox").
				Columns("id", "event", "created_at").
				Values(
					e.EventId(),
					e.Proto(), //todo proto marshall into []byte
					time.Now(),
				)

			if _, err := query.RunWith(o.conn(ctx)).ExecContext(ctx); err != nil {
				return fmt.Errorf("error add event into transactional outbox. %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error execute transaction. %w", err)
	}

	return nil
}
