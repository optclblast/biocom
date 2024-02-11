package transactionaloutbox

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/events"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/logger"
	"github.com/optclblast/biocom/pkg/sqlutils"
)

type transactionalOutboxCron struct {
	mu               sync.Mutex
	log              *slog.Logger
	isRunning        bool
	fn               func()
	eventsInteractor events.EventsInteractor
	db               *sql.DB
}

func NewTransactionalOutboxCron(eventsInteractor events.EventsInteractor, db *sql.DB, log *slog.Logger) *transactionalOutboxCron {
	cron := &transactionalOutboxCron{
		log:              log,
		eventsInteractor: eventsInteractor,
		db:               db,
	}

	cron.fn = cron.sendEventsCron()

	return cron
}

func (c *transactionalOutboxCron) Run() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isRunning {
		return
	}

	go c.fn()
	c.isRunning = true
}

const period time.Duration = 2 * time.Second

func (c *transactionalOutboxCron) sendEventsCron() func() {
	return func() {
		log := c.log.With(
			slog.StringValue("transactional_outbox_cron"),
		)

		timer := time.NewTicker(period)

		for {
			func() {
				txctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
				defer cancel()

				err := c.transaction(txctx, func(ctx context.Context) error {
					var err error

					selectQuery := squirrel.Select("id", "event").
						From("transactional_outbox").
						OrderBy("created_at asc").
						Limit(5)

					rows, err := selectQuery.RunWith(c.conn(ctx)).QueryContext(ctx)
					if err != nil {
						return fmt.Errorf("error fetch evets from database. %w", err)
					}
					defer func() {
						if closeErr := rows.Close(); closeErr != nil {
							err = errors.Join(fmt.Errorf("error close rows. %w", closeErr), err)
						}
					}()

					events := make([]events.Event, 0, 5)
					ids := make([]uuid.UUID, 0, 5)

					for rows.Next() {
						var id uuid.UUID
						var event []byte

						if err := rows.Scan(&id, &event); err != nil {
							return fmt.Errorf("error scan rows. %w", err)
						}

						ids = append(ids, id)
					}

					if len(events) == 0 {
						return nil
					}

					deleteQuery := squirrel.Delete("transactional_outbox").
						Where("id = ?", ids)

					if _, err := deleteQuery.RunWith(c.conn(ctx)).ExecContext(ctx); err != nil {
						return fmt.Errorf("error delete events from transactionsl outbox. %w", err)
					}

					err = c.eventsInteractor.Publish(ctx, events...)
					if err != nil {
						return fmt.Errorf("error publish event. %w", err)
					}

					return err
				})
				if err != nil {
					log.Error("error execute transaction", logger.Err(err))
				}

				<-timer.C
				timer.Reset(period)
			}()
		}
	}
}

func (s *transactionalOutboxCron) conn(ctx context.Context) sqlutils.DBTX {
	if tx, ok := ctx.Value(transactionCtxKey{}).(*sql.Tx); ok {
		return tx
	}

	return s.db
}

func (s *transactionalOutboxCron) transaction(ctx context.Context, f func(ctx context.Context) error) error {
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

	// todo backoff

	err = f(ctx)
	if err != nil {
		return fmt.Errorf("error run transaction function: %w", err)
	}

	return err
}
