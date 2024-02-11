package mappers

import (
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/models"
	userv1 "github.com/optclblast/biocom/pkg/proto/gen/warden/user/v1"
)

func MapUsersToProto(users []*models.User) []*userv1.User {
	var usersProto []*userv1.User = make([]*userv1.User, len(users))
	for i, user := range users {
		usersProto[i] = &userv1.User{
			Id:        user.Id,
			Login:     user.Login,
			Name:      user.Name,
			CreatedAt: uint64(user.CreatedAt.UnixMilli()),
			UpdatedAt: uint64(user.UpdatedAt.UnixMilli()),
			DeletedAt: uint64(user.DeletedAt.UnixMilli()),
		}
	}

	return usersProto
}
