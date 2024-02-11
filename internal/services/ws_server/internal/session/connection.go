package session

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/usecase/controllers"
	"github.com/optclblast/biocom/internal/services/ws_server/pkg/reqctx"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

const connectionLifetime = 15 * time.Minute

type Connection struct {
	mu      *sync.RWMutex
	log     *slog.Logger
	id      uuid.UUID
	conn    *websocket.Conn
	closing bool
	pts     atomic.Uint64
	timer   *time.Ticker

	ctx        context.Context
	reqctx     reqctx.RequestContext
	ctxPool    reqctx.RequestContextPool
	controller *controllers.RootController
}

func (c *Connection) SetLogger(l *slog.Logger) {
	c.log = l
}

func (c *Connection) Id() uuid.UUID {
	return c.id
}

func (c *Connection) Close() error {
	c.mu.Lock()
	if c.closing {
		c.mu.Unlock()
		return nil
	}

	c.closing = true
	c.mu.Unlock()

	return c.conn.Close()
}

func (c *Connection) AttachContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *Connection) Read() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if c.closing {
				return
			}

			ctx, cancel := context.WithTimeout(c.ctx, time.Second*5)
			defer cancel()

			err := (func(ctx context.Context) error {
				_, data, err := c.conn.ReadMessage()
				if err != nil {
					return fmt.Errorf("error read from connection. %w", err)
				}

				err = c.Dispatch(data)
				if err != nil {
					return fmt.Errorf("error handle request. %w", err)
				}

				return nil
			})(ctx)

			if err != nil {

				if websocket.IsUnexpectedCloseError(errors.Unwrap(err)) {
					//logger.Warn("conncection %s closed by client", c.id.String())
					return
				}

				//logger.Error(err)
				return
			}
		}
	}
}

func (c *Connection) Dispatch(b []byte) error {
	c.timer.Reset(connectionLifetime)

	var request *apiv1.Request = new(apiv1.Request)
	err := proto.Unmarshal(b, request)
	if err != nil {
		return fmt.Errorf("error unmarshal proto request message. %w", err)
	}

	ctx := c.newRequestContext(request, c.ctx)
	defer func() {
		c.freeRequestContext()
	}()

	err = c.dispatchRequest(ctx)
	if err != nil {
		// send error
	}

	return nil
}

func (c *Connection) Write(resp *apiv1.Response) error {
	out, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("error marshal response. %w", err)
	}

	return c.conn.WriteMessage(websocket.BinaryMessage, out)
}

func (c *Connection) newRequestContext(
	req *apiv1.Request,
	ctx context.Context,
) reqctx.RequestContext {
	return c.ctxPool.Get(req, c.pts.Load(), ctx)
}

func (c *Connection) freeRequestContext() {
	c.ctxPool.Put(c.reqctx)
}

func (c *Connection) dispatchRequest(req reqctx.RequestContext) error {
	switch {
	case req.Request().GetAuthSignIn() != nil:
		return c.handle(req, c.controller.AuthController.SignIn, "sign_in")
	case req.Request().GetAuthSignUp() != nil:
		return c.handle(req, c.controller.AuthController.SignUp, "sign_up")
	default:
		return nil // todo return error
	}
}

func (c *Connection) handle(
	req reqctx.RequestContext,
	handler func(req reqctx.RequestContext) (*apiv1.Response, error),
	handlerName string,
) error {
	resp, err := handler(req)
	if err != nil {
		return fmt.Errorf("error handle request. %w", err)
	}

	return c.Write(resp)
}
