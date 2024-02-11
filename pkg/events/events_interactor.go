package events

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	eventsv1 "github.com/optclblast/biocom/pkg/proto/gen/events"
	"google.golang.org/protobuf/proto"
)

type EventsInteractor interface {
	Publish(ctx context.Context, events ...Event) error
}

func NewEventsInteractor(publisher jetstream.Publisher) EventsInteractor {
	return &eventsInteractorNats{
		publisher: publisher,
	}
}

type eventsInteractorNats struct {
	publisher jetstream.Publisher
}

func (e *eventsInteractorNats) Publish(ctx context.Context, events ...Event) error {
	subjectEvents := make(map[string][]*eventsv1.Event)

	for _, event := range events {
		if batch, ok := subjectEvents[event.Subject()]; ok {
			batch = append(batch, event.Proto())
			continue
		}

		subjectEvents[event.Subject()] = []*eventsv1.Event{event.Proto()}
	}

	for subject, protoEvents := range subjectEvents {
		protoMessage := &eventsv1.Events{
			ServerTime: uint64(time.Now().UnixMilli()),
			// CompanyId:  companyId,
			Events: protoEvents,
		}

		data, err := proto.Marshal(protoMessage)
		if err != nil {
			return fmt.Errorf("error marshal event. %w", err)
		}

		_, err = e.publisher.Publish(ctx, subject, data)
		if err != nil {
			return fmt.Errorf("error publish events. %w", err)
		}
	}

	return nil
}
