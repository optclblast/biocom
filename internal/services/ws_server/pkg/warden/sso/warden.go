package sso

import (
	"context"
	"fmt"

	ssov1 "github.com/optclblast/biocom/pkg/proto/gen/warden/sso/v1"
	"google.golang.org/grpc"
)

type WardenSSOClient interface {
	SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error)
	SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error)
}

type NewWardenSSOParams struct {
	Target string
}

func NewWardenSSO(params NewWardenSSOParams) (WardenSSOClient, func(), error) {
	grpcClient, err := grpc.Dial(params.Target)
	if err != nil {
		return nil, func() {}, fmt.Errorf("error connect to WardenSSO api. %w", err)
	}

	return &wardenSSOClient{
		client: ssov1.NewWardenSSOAPIClient(grpcClient),
	}, func() { grpcClient.Close() }, nil
}

type wardenSSOClient struct {
	client ssov1.WardenSSOAPIClient
}

func (w *wardenSSOClient) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	return w.client.SignIn(ctx, req)
}

func (w *wardenSSOClient) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {
	return w.client.SignUp(ctx, req)
}
