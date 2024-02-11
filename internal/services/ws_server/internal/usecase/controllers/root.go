package controllers

import (
	"github.com/optclblast/biocom/internal/services/ws_server/pkg/reqctx"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

type Controller interface {
	Invoke(req reqctx.RequestContext) (*apiv1.Response, error)
}

type RootController struct {
	AuthController AuthController
}

func NewRootController(
	authController AuthController,
) *RootController {
	return &RootController{
		AuthController: authController,
	}
}
