package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WardenGRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

type GRPCService interface {
	Register(server *grpc.Server)
}

func New(
	log *slog.Logger,
	port int,
	services ...GRPCService,
) *WardenGRPC {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))

	for _, service := range services {
		service.Register(gRPCServer)
	}

	return &WardenGRPC{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// InterceptorLogger adapts slog logger to interceptor logger.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (a *WardenGRPC) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *WardenGRPC) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("error listen tcp at %v. %w", a.port, err)
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("error serve grpc. %w", err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *WardenGRPC) Stop() {
	a.log.Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
