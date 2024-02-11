package service

import (
	"log/slog"

	srvconf "github.com/optclblast/biocom/internal/services/warden/internal/config"
	"github.com/optclblast/biocom/internal/services/warden/internal/service/grpc"
	"go-micro.dev/v4/config"
)

type Warden struct {
	gRPCServer *grpc.WardenGRPC
	config     config.Config
}

func New(
	config config.Config,
	log *slog.Logger,
	services ...grpc.GRPCService,
) *Warden {
	conf, err := srvconf.MapConfig(config.Bytes())
	if err != nil {
		panic(err)
	}

	log.Debug("starting warden service", slog.Int("port", conf.GRPCPort))

	return &Warden{
		gRPCServer: grpc.New(log, conf.GRPCPort, services...),
		config:     config,
	}
}

func (w *Warden) Run() {
	go w.gRPCServer.MustRun()
}

func (a *Warden) Stop() {
	a.gRPCServer.Stop()
}
