package nats

import (
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
)

type publisherConsumer interface {
	jetstream.Publisher
	jetstream.Consumer
}

type natsJobManager struct {
	js         publisherConsumer
	numWorkers uint8 // up to 255 active workers
	log        *slog.Logger
	//pre
}
