package nomenclature

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/optclblast/biocom/internal/services/garganta/internal/models"
	repo "github.com/optclblast/biocom/internal/services/garganta/internal/usecase/repository/nomenclature"
)

type repositorySQL struct {
	db *sql.DB
}

func NewRepositorySQL(db *sql.DB) repo.Repository {
	return &repositorySQL{
		db: db,
	}
}

func (r *repositorySQL) Nomenclature(ctx context.Context, params repo.NomenclatureParams) (models.Item, error) {
	query := sq.Select("nomenclature").Columns("id, company_id, name, created_at, updated_at, type").
		Where(sq.Eq{"company_id": params.CompanyId})

	if params.Id != "" {
		query = query.Where(sq.Eq{"id": params.Id})
	}
	if !params.WithDeleted {
		query = query.Where(sq.Eq{"deleted_at": nil})
	}

	if params.WithComposition {
		// TODO fetch nomenclature composition
	}

	rows, err := query.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetch nomenclature from database. %w", err)
	}

	for rows.Next() {
		var (
			id               string
			companyId        string
			name             string
			createdAt        time.Time
			updatedAt        time.Time
			nomenclatureType models.NomenclatureType
		)

		if err = rows.Scan(
			&id,
			&companyId,
			&name,
			&createdAt,
			&updatedAt,
			&nomenclatureType,
		); err != nil {
			return nil, fmt.Errorf("error scan row. %w", err)
		}

		return models.NewNomenclature(
			id,
			companyId,
			name,
			createdAt,
			updatedAt,
			time.Time{},
			nomenclatureType,
		)
	}

	return nil, ErrNotFound
}

func (r *repositorySQL) ListNomenclature(ctx context.Context, params repo.ListNomenclatureParams) ([]models.Item, error) {
	items := make([]models.Item, len(params.Ids))

	for i, id := range params.Ids {
		item, err := r.Nomenclature(ctx, repo.NomenclatureParams{
			Id:          id,
			CompanyId:   params.CompanyId,
			WithDeleted: params.WithDeleted,
		})
		if err != nil {
			return nil, fmt.Errorf("erro fetch nomenclature from database. %w", err)
		}

		items[i] = item
	}

	return items, nil
}

func (r *repositorySQL) UpdateNomenclature(ctx context.Context, params repo.UpdateNomenclatureParams) error {
	//todo update logic
	return nil
}

func (r *repositorySQL) DeleteNomenclature(ctx context.Context, params repo.DeleteNomenclatureParams) error {
	query := sq.Delete("nomenclature").Where(sq.Eq{"id": params.Id})
	if _, err := query.RunWith(r.db).ExecContext(ctx); err != nil {
		return fmt.Errorf("error delete nomenclature from database. %w", err)
	}
	return nil
}
