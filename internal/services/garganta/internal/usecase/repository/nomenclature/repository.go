package nomenclature

import (
	"context"

	"github.com/optclblast/biocom/internal/services/garganta/internal/models"
)

type NomenclatureParams struct {
	Id          string
	CompanyId   string
	WithDeleted bool
}

type ListNomenclatureParams struct {
	Ids         []string
	CompanyId   string
	WithDeleted bool
}

type UpdateNomenclatureParams struct {
	Id        string
	CompanyId string
	// todo
}

type DeleteNomenclatureParams struct {
	Id string
}

type Repository interface {
	Nomenclature(ctx context.Context, params NomenclatureParams) (models.Item, error)
	ListNomenclature(ctx context.Context, params ListNomenclatureParams) ([]models.Item, error)
	UpdateNomenclature(ctx context.Context, params UpdateNomenclatureParams) error
	DeleteNomenclature(ctx context.Context, params DeleteNomenclatureParams) error
}
