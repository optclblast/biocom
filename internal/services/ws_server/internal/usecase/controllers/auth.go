package controllers

import (
	"fmt"

	"github.com/optclblast/biocom/internal/services/ws_server/pkg/reqctx"
	ssov1 "github.com/optclblast/biocom/pkg/proto/gen/warden/sso/v1"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

type AuthController interface {
	SignIn(req reqctx.RequestContext) (*apiv1.Response, error)
	SignUp(req reqctx.RequestContext) (*apiv1.Response, error)
}

type AuthControllerImpl struct {
	wardenSSO ssov1.WardenSSOAPIClient
}

func NewAuthController() AuthController {
	return &AuthControllerImpl{}
}

func (c *AuthControllerImpl) SignIn(req reqctx.RequestContext) (*apiv1.Response, error) {
	request := req.Request().GetAuthSignIn()

	_, err := c.wardenSSO.SignIn(req.Context(), &ssov1.SignInRequest{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("error authorize user. %w", err)
	}

	// todo вот тут пока непонятно, как сделать

	return &apiv1.Response{
		Id:      req.Request().GetId(),
		Payload: &apiv1.Response_AuthSignIn{},
	}, nil
}

func (c *AuthControllerImpl) SignUp(req reqctx.RequestContext) (*apiv1.Response, error) {
	return nil, nil
}
