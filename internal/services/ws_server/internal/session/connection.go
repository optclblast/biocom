package session

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/usecase/controllers"
	"github.com/optclblast/biocom/internal/services/ws_server/pkg/reqctx"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

const connectionLifetime = 15 * time.Minute

type Connection struct {
	mu         *sync.RWMutex
	log        *slog.Logger
	id         string
	conn       *websocket.Conn
	sesStorage SessionPool
	ctxPool    reqctx.RequestContextPool

	closing bool
	timer   *time.Ticker

	ctx        context.Context
	controller *controllers.RootController
}

func (c *Connection) SetLogger(l *slog.Logger) {
	c.log = l
}

func (c *Connection) Id() string {
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

				err = c.dispatch(data)
				if err != nil {
					return fmt.Errorf("error handle request. %w", err)
				}

				return nil
			})(ctx)

			if err != nil {

				if websocket.IsUnexpectedCloseError(errors.Unwrap(err)) {
					c.log.Warn("conncection closed by client", slog.String("id", c.id))
					c.Close()
					return
				}

				c.log.Error("error handle ws message", slog.String("error", err.Error()))
				return
			}
		}
	}
}

func (c *Connection) dispatch(b []byte) error {
	var err error

	c.timer.Reset(connectionLifetime)

	var request *apiv1.Request = new(apiv1.Request)
	err = proto.Unmarshal(b, request)
	if err != nil {
		return fmt.Errorf("error unmarshal proto request message. %w", err)
	}

	var session Session
	switch {
	case request.GetSessionInit() != nil:
		session = c.sessionInit(request)
	default:
		claims := make(jwt.MapClaims)
		_, err := jwt.ParseWithClaims(request.GetToken(), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("REAL_SECRET"), nil // TODO
		})
		if err != nil {
			return fmt.Errorf("error parse token. %w", err)
		}

		if sessionId, ok := claims["session_id"].(string); ok {
			session = c.sesStorage.Get(sessionId)
		} else {
			session = c.sesStorage.Get("")
		}
	}

	ctx := c.ctxPool.Get(request, session.Pts(), c.ctx)
	return session.DispatchRequest(ctx)
}

func (c *Connection) Write(resp *apiv1.Response) error {
	out, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("error marshal response. %w", err)
	}

	return c.conn.WriteMessage(websocket.BinaryMessage, out)
}

func (c *Connection) sessionInit(req *apiv1.Request) Session {
	sesInitReq := req.GetSessionInit()

	var sessionId string = sesInitReq.GetSessionId()

	if sesInitReq.GetSessionId() == "" {
		sessionId = uuid.NewString()
	}

	return c.sesStorage.Get(sessionId)
}
