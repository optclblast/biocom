package reqctx

import (
	"context"
	"sync"
	"time"

	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

type RequestContextPool interface {
	Get(
		req *apiv1.Request,
		pts uint64,
		ctx context.Context,
	) RequestContext
	Put(RequestContext)
}

type requestContextPool struct {
	p sync.Pool
}

func NewRequestContextPool() RequestContextPool {
	return &requestContextPool{
		p: sync.Pool{New: func() any {
			return new(requestCtx)
		}},
	}
}

func (p *requestContextPool) Get(
	req *apiv1.Request,
	pts uint64,
	ctx context.Context,
) RequestContext {
	return p.p.Get().(*requestCtx)
}

func (p *requestContextPool) Put(v RequestContext) {
	p.p.Put(v)
}

// todo
type RequestContext interface {
	context.Context

	Pts() uint64
	Set(k any, v any)
	Value(k any) any
	Request() *apiv1.Request
	Context() context.Context
}

type requestCtx struct {
	ctx context.Context

	pts      uint64
	done     chan struct{}
	deadline time.Time
	request  *apiv1.Request
}

func (c *requestCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, false //todo
}

func (c *requestCtx) Done() <-chan struct{} {
	return c.done //todo
}

func (c *requestCtx) Err() error {
	return nil //todo
}

func (c *requestCtx) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *requestCtx) Pts() uint64 {
	return c.pts
}

func (c *requestCtx) Set(k any, v any) {
	c.ctx = context.WithValue(c.ctx, k, v)
}

func (c *requestCtx) Request() *apiv1.Request {
	return c.request
}

func (c *requestCtx) Context() context.Context {
	return c.ctx
}
