package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/optclblast/biocom/pkg/jobs"
	"github.com/optclblast/biocom/pkg/sqlutils"
)

type postgresJobQueue struct {
	db *sql.DB
}

func NewPostgresJobQueue(db *sql.DB) jobs.JobQueue {
	return &postgresJobQueue{db}
}

func (p *postgresJobQueue) Push(job jobs.Job) error {
	jobData, err := job.Marshal()
	if err != nil {
		return fmt.Errorf("error marshal job. %w", err)
	}

	query := sq.Insert("jobs").
		Columns("id", "job", "priority", "ttr").
		Values(
			job.JobId(),
			jobData,
			job.Priority(),
			job.TimeToRun(),
		).Suffix("on conflict (id) do nothing")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = query.ExecContext(ctx); err != nil {
		return fmt.Errorf("error add job into a queue. %w", err)
	}

	return nil
}

func (p *postgresJobQueue) Pop() (jobs.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var job jobs.Job

	err := p.Transaction(ctx, func(ctx context.Context) error {
		nextJobQuery := sq.Select("jobs").Column("id", "job", "ttr").
			Where(`status = 0 AND delayed_to >= now()`).
			OrderBy("proprity DESC").Limit(1).Suffix("for update skip locked")

		row := nextJobQuery.QueryRowContext(ctx)

		var (
			id  string
			job []byte
			ttr int64
		)

		if err := row.Scan(&id, &job, &ttr); err != nil {
			return fmt.Errorf("error scan row. %w", err)
		}

		// todo build job

		updateJobStatusQuery := sq.Update("jobs").Set("status", 1)

		if _, err := updateJobStatusQuery.ExecContext(ctx); err != nil {
			return fmt.Errorf("error update job status. %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error fetch job from the queue. %w", err)
	}

	return job, nil
}

func (p *postgresJobQueue) Close() {
	p.db.Close()
}

type transactionCtxKey struct{}

func (s *postgresJobQueue) Conn(ctx context.Context) sqlutils.DBTX {
	if tx, ok := ctx.Value(transactionCtxKey{}).(*sql.Tx); ok {
		return tx
	}

	return s.db
}

func (s *postgresJobQueue) Transaction(
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
