package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	consulconf "github.com/go-micro/plugins/v4/config/source/consul"
	"github.com/google/uuid"
	"github.com/optclblast/biocom/internal/services/ws_server/cmd/commands"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/server"
	"github.com/optclblast/biocom/pkg/logger"
	consulReg "github.com/optclblast/biocom/pkg/registry/consul"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
)

func main() {
	app := &cli.App{
		Name:     "openerp-ws-server",
		Version:  "0.0.1a",
		Commands: commands.Commands(),
		Flags:    []cli.Flag{},
		Action: func(c *cli.Context) error {
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

			consulSource := consulconf.NewSource(
				consulconf.WithAddress("localhost:8500"),
				consulconf.WithPrefix("warden/"),
				consulconf.StripPrefix(true),
			)

			registry, err := consulReg.NewConsul("localhost:8500", "dc1", "")
			if err != nil {
				return fmt.Errorf("error create new consul registry client. %w", err)
			}

			_, err = config.NewConfig(config.WithSource(consulSource))
			if err != nil {
				return fmt.Errorf("error fetch config from consul. %w", err)
			}

			log := logger.NewLogger(slog.LevelDebug)

			server := server.NewWebsocketServer(
				server.WithLogger(log),
			)

			log.Info("server started", slog.String("address", server.Address()))

			registry.Register(c.App.Name, uuid.NewString(), server.Address())
			log.Info("service registered")

			defer func() {
				log.Info("leaving node")
				registry.Unregister(c.App.Name, uuid.NewString())
			}()

			if err := server.Run(c.Context); err != nil {
				return fmt.Errorf("error run application. %w", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
