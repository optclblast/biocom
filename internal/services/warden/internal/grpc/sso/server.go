package auth

import (
	"context"

	"github.com/optclblast/biocom/internal/services/warden/internal/controller"
	service "github.com/optclblast/biocom/internal/services/warden/internal/service/grpc"
	ssov1 "github.com/optclblast/biocom/pkg/proto/gen/warden/sso/v1"
	"google.golang.org/grpc"
)

type WardenAuthAPI interface {
	service.GRPCService
	SignIn(context.Context, *ssov1.SignInRequest) (*ssov1.SignInResponse, error)
	SignUp(context.Context, *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error)
}

type authService struct {
	*ssov1.UnimplementedWardenSSOAPIServer

	authController *controller.AuthController
}

func NewWardenAuthService(
	authController *controller.AuthController,
) WardenAuthAPI {
	return &authService{
		authController: authController,
	}
}

func (s *authService) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	return s.authController.SignIn(ctx, req)
}

func (s *authService) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {
	return s.authController.SignUp(ctx, req)
}

func (s *authService) Register(server *grpc.Server) {
	ssov1.RegisterWardenSSOAPIServer(server, s)
}
