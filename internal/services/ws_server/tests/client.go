package tests

import (
	"net/url"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

type WSDoer interface {
	Do(data []byte) ([]byte, error)
	Close() error
}

type TestClient struct {
	pts          atomic.Uint64
	login        string
	password     string
	onganization string
	conn         *websocket.Conn
}

func NewTestClient(login, pass, org, addr string) WSDoer {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	return &TestClient{
		pts:          atomic.Uint64{},
		login:        login,
		password:     pass,
		onganization: org,
		conn:         conn,
	}
}

func (c *TestClient) Do(data []byte) ([]byte, error) {
	err := c.conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return nil, err
	}

	_, re, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	return re, nil
}

func (c *TestClient) Close() error {
	return c.conn.Close()
}
