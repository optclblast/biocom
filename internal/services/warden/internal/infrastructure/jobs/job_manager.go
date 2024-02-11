package jobs

import (
	"context"
	"time"
)

type Manager interface {
	RegisterJob(Job, JobUnmashaler)
	AddJob(Job) (id string, err error)
	Run() error
	Shutdown() error
}

type Job interface {
	JobName() string
	JobId() string
	Priority() int
	Delay() time.Duration
	TimeToRun() time.Duration
	Execute(ctx context.Context) error

	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type JobUnmashaler func(b []byte) (Job, error)

type JobQueue interface {
	Add(i Job) error
	Remove() (Job, error)
	Close()
	CloseRemaining() []Job
	Closed() bool
	Wait() (Job, error)
	Cap() int
	Len() int
}
