package events

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	eventsv1 "github.com/optclblast/biocom/pkg/proto/gen/events"
)

type EventsBuilder interface {
	UserSignIn(userId string, companyId string, time time.Time) Event
	UserSignUp(userId string, companyId string, time time.Time) Event
}

func NewEventsBuilder(config BrokerConfig) EventsBuilder {
	return &eventsBuilderNats{config.Datacenter()}
}

type eventsBuilderNats struct {
	dc string // datacenter
}

type Event interface {
	EventId() uuid.UUID
	Proto() *eventsv1.Event
	Subject() string
}

// Event object
type event struct {
	id      uuid.UUID
	proto   *eventsv1.Event
	subject string
}

func (e *event) EventId() uuid.UUID {
	return e.id
}

func (e *event) Proto() *eventsv1.Event {
	return e.proto
}

func (e *event) Subject() string {
	return e.subject
}

// Nats subjects to push in {

// <DC>.warden.log subject. Subject that holds all important events that will be logged in log-service
func (e *eventsBuilderNats) subjectWardenLog() string {
	return fmt.Sprintf("%s.warden.log", e.dc)
}

// ,DC..warden.user subject. This subject holds all event that related to a specific user,
// like name changed, avatar changes, position changed, etc
func (e *eventsBuilderNats) subjectWardenUser() string {
	return fmt.Sprintf("%s.warden.user", e.dc)
}

func (e *eventsBuilderNats) subjectWardenCompany() string {
	return fmt.Sprintf("%s.warden.company", e.dc)
}

// }

func (e *eventsBuilderNats) UserSignIn(userId string, companyId string, time time.Time) Event {
	eventId := uuid.New()

	return &event{
		id: eventId,
		proto: &eventsv1.Event{
			IdempotencyKey: fmt.Sprintf("%s%s%s", userId, companyId, eventId.String()),
			Payload: &eventsv1.Event_Login{
				Login: &eventsv1.EventUserLogin{
					UserId:    userId,
					CompanyId: companyId,
					Time:      uint64(time.UnixMilli()),
				},
			},
		},
		subject: e.subjectWardenLog(),
	}
}

func (e *eventsBuilderNats) UserSignUp(userId string, companyId string, time time.Time) Event {
	eventId := uuid.New()

	return &event{
		id: eventId,
		proto: &eventsv1.Event{
			IdempotencyKey: fmt.Sprintf("%s%s%s", userId, companyId, eventId.String()),
			Payload: &eventsv1.Event_SignUp{
				SignUp: &eventsv1.EventUserSignedUp{
					UserId:    userId,
					CompanyId: companyId,
					Time:      uint64(time.UnixMilli()),
				},
			},
		},
		subject: e.subjectWardenLog(),
	}
}
