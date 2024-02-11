package controller

import (
	"context"
	"fmt"

	"github.com/optclblast/biocom/internal/services/warden/internal/lib/mappers"
	"github.com/optclblast/biocom/internal/services/warden/internal/usecase/repository/user"
	userv1 "github.com/optclblast/biocom/pkg/proto/gen/warden/user/v1"
)

type UsersController struct {
	usersRepository user.Repository
}

func (c *UsersController) GetUsers(
	ctx context.Context,
	req *userv1.GetUsersRequest,
) (*userv1.GetUsersResponse, error) {
	if len(req.GetIds()) == 0 && req.GetOrganizationId() == "" {
		return nil, fmt.Errorf("error empty request")
	}

	users, err := c.usersRepository.GetUsers(ctx, user.GetUsersParams{
		Ids:            req.GetIds(),
		OrganizationId: req.GetOrganizationId(),
	})
	if err != nil {
		return nil, fmt.Errorf("error fetch users from repository. %w", err)
	}

	return &userv1.GetUsersResponse{
		Users: mappers.MapUsersToProto(users),
	}, nil
}
