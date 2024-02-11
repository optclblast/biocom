package nats

import (
	"context"
	"fmt"

	"time"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/jobs"
)

type natsJobQueue struct {
	queueName string
	js        jetstream.JetStream
}

type natsJob struct{}

func (n *natsJobQueue) Add(job jobs.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := job.Marshal()
	if err != nil {
		return fmt.Errorf("error marshal job. %w", err)
	}

	_, err = n.js.Publish(ctx, n.queueName, data)
	if err != nil {
		return fmt.Errorf("error publish job into job queue subject. %w", err)
	}

	return nil
}

// todo implement
func (n *natsJobQueue) Remove() (jobs.Job, error)  { return nil, nil }
func (n *natsJobQueue) Close()                     {}
func (n *natsJobQueue) CloseRemaining() []jobs.Job { return nil }
func (n *natsJobQueue) Closed() bool               { return true }
func (n *natsJobQueue) Wait() (jobs.Job, error)    { return nil, nil }
func (n *natsJobQueue) Cap() int                   { return 0 }
func (n *natsJobQueue) Len() int                   { return 0 }

func NewNatsJobQueue(queueName string, js jetstream.JetStream) jobs.JobQueue {
	return &natsJobQueue{
		queueName: queueName,
		js:        js,
	}
}
