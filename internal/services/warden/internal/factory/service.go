package factory

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	srvconf "github.com/optclblast/biocom/internal/services/warden/internal/config"
	"github.com/optclblast/biocom/internal/services/warden/internal/controller"
	auth "github.com/optclblast/biocom/internal/services/warden/internal/grpc/sso"
	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/events"
	usersql "github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/repository/user"
	transactionaloutbox "github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/transactional_outbox"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/logger"
	"github.com/optclblast/biocom/internal/services/warden/internal/service"
	"github.com/optclblast/biocom/internal/services/warden/internal/usecase/repository/user"
	"go-micro.dev/v4/config"
)

// ProvideWardenService returns Warden service and cleanup function or panics
func ProvideWardenService(config config.Config) (*service.Warden, func()) {
	cfg := provideConfig(config)

	userRepo, cleanup, err := provideUserRepository(cfg.SSODB)
	if err != nil {
		panic(fmt.Errorf("error provide user repository. %w", err))
	}

	log := logger.NewLogger(logger.MapLevel(cfg.Env))

	log.Debug("CONFIG", slog.AnyValue(cfg))

	nc, err := nats.Connect(cfg.Broker.Address)
	if err != nil {
		panic(fmt.Errorf("error connecto to nats. %w", err))
	}

	natsJs, err := jetstream.New(nc)
	if err != nil {
		panic(fmt.Errorf("error connecto to jetstream. %w", err))
	}

	authApi := auth.NewWardenAuthService(
		controller.NewAuthController(
			log,
			15*time.Minute,
			userRepo,
			transactionaloutbox.NewTransactionalOutbox(
				provideTransactionalOutboxDB(),
				log,
				events.NewEventsInteractor(natsJs),
			),
			events.NewEventsBuilder(*cfg.Broker),
		),
	)

	warden := service.New(config, log, authApi)

	return warden, func() {
		cleanup()
	}
}

func provideUserRepository(config *srvconf.DatabaseConfig) (user.Repository, func(), error) {
	conn, err := sql.Open(
		"postgres", fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User,
			config.Password, config.DBName, config.SSLMode,
		),
	)
	if err != nil {
		return nil, func() {}, fmt.Errorf("error connect to a database. %w", err)
	}

	return usersql.NewAuthSQL(conn), func() {
		conn.Close()
	}, nil
}

func provideConfig(cfg config.Config) *srvconf.Config {
	conf, err := srvconf.MapConfig(cfg.Bytes())
	if err != nil {
		panic(err)
	}

	return conf
}

func provideTransactionalOutboxDB() *sql.DB {
	return nil // TODO
}
