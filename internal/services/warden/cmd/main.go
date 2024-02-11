package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	consulconf "github.com/go-micro/plugins/v4/config/source/consul"

	//"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/optclblast/biocom/internal/services/warden/cmd/commands"
	"github.com/optclblast/biocom/internal/services/warden/internal/factory"

	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
)

func main() {
	app := &cli.App{
		Name:     "biocom",
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

			conf, err := config.NewConfig(config.WithSource(consulSource))
			if err != nil {
				return fmt.Errorf("error fetch config from consul. %w", err)
			}

			warden, cleanup := factory.ProvideWardenService(conf)
			defer cleanup()

			warden.Run()

			<-stop
			warden.Stop()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
