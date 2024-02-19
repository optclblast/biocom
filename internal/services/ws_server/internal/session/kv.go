package session

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
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

func (n *NatsSessionStorage) Get(ctx context.Context, id string) (Session, error) {
	entry, err := n.kv.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error fetch session from bucket. %w", err)
	}

	sessionProto := new(apiv1.Session)

	if err := proto.Unmarshal(entry.Value(), sessionProto); err != nil {
		return nil, fmt.Errorf("error unmarshal session from raw data. %w", err)
	}

	return FromProto(sessionProto), nil
}

func (n *NatsSessionStorage) Add(ctx context.Context, session Session) error {
	data := session.ToProto()

	_, err := n.kv.Put(ctx, session.Id(), data)
	if err != nil {
		return fmt.Errorf("error put session into storage. %w", err)
	}

	return nil
}

func (n *NatsSessionStorage) Update(ctx context.Context, session Session) error {
	data := session.ToProto()

	_, err := n.kv.Update(ctx, session.Id(), data, uint64(time.Now().UnixMilli()))
	if err != nil {
		return fmt.Errorf("error put session into storage. %w", err)
	}

	return nil
}
