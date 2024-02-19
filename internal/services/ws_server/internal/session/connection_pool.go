package session

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionPool struct {
	mu    *sync.RWMutex
	p     sync.Pool
	conns map[string]*Connection
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		p: sync.Pool{
			New: func() any {
				return new(Connection)
			},
		},
		mu:    &sync.RWMutex{},
		conns: make(map[string]*Connection),
	}
}

func (c *ConnectionPool) Get(wsconn *websocket.Conn) *Connection {
	conn := c.p.Get().(*Connection)

	conn.id = uuid.NewString()
	conn.mu = &sync.RWMutex{}
	conn.conn = wsconn
	conn.ctx = context.TODO()

	(func() {
		c.mu.Lock()

		defer func() {
			c.mu.Unlock()
		}()

		c.conns[conn.id] = conn
	}())

	return conn
}

func (c *ConnectionPool) Put(conn *Connection) {
	delete(c.conns, conn.Id())

	c.p.Put(conn)
}

func (c *ConnectionPool) Connections() (map[string]*Connection, func()) {
	c.mu.Lock()
	return c.conns, func() {
		c.mu.Unlock()
	}
}
