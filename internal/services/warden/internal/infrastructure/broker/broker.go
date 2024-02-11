package broker

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/optclblast/biocom/internal/services/warden/internal/config"
)

func NewJetstream(config *config.BrokerConfig) (jetstream.JetStream, error) {
	nc, err := nats.Connect(config.Address)
	if err != nil {
		return nil, fmt.Errorf("error connect to nats server. %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("error create new jetstream instance. %w", err)
	}

	return js, nil
}

func InitWardenStream(
	ctx context.Context,
	js jetstream.JetStream,
	subjects ...string,
) (jetstream.Stream, error) {
	wardenStream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "warden_stream",
		Subjects: subjects,
	})
	if err != nil {
		if !errors.Is(err, jetstream.ErrStreamNameAlreadyInUse) {
			return nil, fmt.Errorf("error create warden stream. %w", err)
		}
	}

	return wardenStream, nil
}
