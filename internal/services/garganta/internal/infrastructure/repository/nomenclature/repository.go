package nomenclature

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
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
	query := squirrel.Select("nomenclature").Columns("id")
}

func (r *repositorySQL) ListNomenclature(ctx context.Context, params repo.ListNomenclatureParams) ([]models.Item, error) {

}

func (r *repositorySQL) UpdateNomenclature(ctx context.Context, params repo.UpdateNomenclatureParams) error {

}

func (r *repositorySQL) DeleteNomenclature(ctx context.Context, params repo.DeleteNomenclatureParams) error {

}
