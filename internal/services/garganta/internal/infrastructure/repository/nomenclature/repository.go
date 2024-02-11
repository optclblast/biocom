package nomenclature

import (
	"database/sql"

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

func (r *repositorySQL) Get()    {}
func (r *repositorySQL) Update() {}
func (r *repositorySQL) Delete() {}
