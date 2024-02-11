package user

import (
	"context"
	"time"

	"github.com/optclblast/biocom/internal/services/warden/internal/lib/models"
)

type CreateParams struct {
	Id             string
	OrganizationId string
	PasswordHash   []byte
	CreatedAt      time.Time
}

type UpdateParams struct {
	Id             string
	OrganizationId string
	UpdatedAt      time.Time
}

type DeleteParams struct {
	Id             string
	OrganizationId string
}

type GetParams struct {
	Id             string
	Login          string
	OrganizationId string
}

type GetUsersParams struct {
	Ids            []string
	OrganizationId string
}

type Repository interface {
	Create(ctx context.Context, params CreateParams) error
	Update(ctx context.Context, params UpdateParams) error
	Delete(ctx context.Context, params DeleteParams) error
	Get(ctx context.Context, params GetParams) (*models.User, error)
	GetUsers(ctx context.Context, params GetUsersParams) ([]*models.User, error)
}
