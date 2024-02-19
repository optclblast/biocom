package jobs

import (
	"time"
)

type Manager interface {
	AddJob(Job) (string, error)
	Run() error
	Shutdown() error
}

type Job interface {
	JobName() string
	JobId() string
	Priority() int
	Delay() time.Duration
	TimeToRun() time.Duration
	ReturnIntoQueue(params ReturnIntoQueueParams) error

	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type JobUnmashaler func(b []byte) (Job, error)

type Queue interface {
	Push(job Job) error
	Pop() (Job, error)
	Close()
}

type ReturnIntoQueueParams struct {
	ErrorCode    int
	ErrorMessage string
	DelayedTo    time.Time
}
