package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/session"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
	"google.golang.org/protobuf/proto"
)

type NatsSessionStorage struct {
	kv jetstream.KeyValue
}

func NewNatsSessionStorage(ctx context.Context, conn *nats.Conn) (*NatsSessionStorage, error) {
	js, err := jetstream.New(conn)
	if err != nil {
		return nil, fmt.Errorf("error create jetstream connection. %w", err)
	}

	kv, err := js.KeyValue(ctx, "sessions")
	if err != nil {
		return nil, fmt.Errorf("error connect to sessions bucket. %w", err)
	}

	return &NatsSessionStorage{
		kv: kv,
	}, nil
}

func (n *NatsSessionStorage) Get(ctx context.Context, id string) (session.Session, error) {
	entry, err := n.kv.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error fetch session from bucket. %w", err)
	}

	sessionProto := new(apiv1.Session)

	if err := proto.Unmarshal(entry.Value(), sessionProto); err != nil {
		return nil, fmt.Errorf("error unmarshal session from raw data. %w", err)
	}

	return session.FromProto(sessionProto), nil
}

func (n *NatsSessionStorage) Add(ctx context.Context, session session.Session) error
func (n *NatsSessionStorage) Update(ctx context.Context, session session.Session) error
