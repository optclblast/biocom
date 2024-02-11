package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/session"
	"github.com/optclblast/biocom/internal/services/ws_server/internal/usecase/controllers"
)

type WebsocketServer struct {
	maxConnectionsLimit uint
	connectionLifetime  uint
	address             string
	log                 *slog.Logger

	connectionPool *session.ConnectionPool

	eventCh        chan any // TODO
	closeCh        chan os.Signal
	rootController *controllers.RootController
}

type websocketServerParams struct {
	maxConnectionsLimit uint
	connectionLifetime  uint
	host                string
	port                int
	logger              *slog.Logger
}

type WSServerOption func(option *websocketServerParams)

func WithMaxConnectionLimit(limit uint) WSServerOption {
	return func(option *websocketServerParams) {
		option.maxConnectionsLimit = limit
	}
}

func WithConnectionLifetime(lt uint) WSServerOption {
	return func(option *websocketServerParams) {
		option.connectionLifetime = lt
	}
}

func WithHost(host string) WSServerOption {
	return func(option *websocketServerParams) {
		option.host = host
	}
}

func WithPort(port int) WSServerOption {
	return func(option *websocketServerParams) {
		option.port = port
	}
}

func WithLogger(logger *slog.Logger) WSServerOption {
	return func(option *websocketServerParams) {
		option.logger = logger
	}
}

func NewWebsocketServer(options ...WSServerOption) *WebsocketServer {
	params := &websocketServerParams{
		maxConnectionsLimit: 500,
		connectionLifetime:  10000,
	}

	for _, opt := range options {
		opt(params)
	}

	params.port = 8080

	if params.port == 0 {
		port, err := freePort()
		if err != nil {
			panic(fmt.Sprintf("error get free port. %s", err.Error()))
		}

		params.port = port
	}

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	controller, err := controllers.InitializeRootController()
	if err != nil {
		log.Fatal("error initialize root controller. ", err.Error())
	}

	return &WebsocketServer{
		maxConnectionsLimit: params.maxConnectionsLimit,
		connectionLifetime:  params.connectionLifetime,
		address:             fmt.Sprintf("%s:%v", params.host, params.port),
		eventCh:             make(chan any), // TODO
		closeCh:             sig,
		connectionPool:      session.NewConnectionPool(),
		log:                 params.logger,
		rootController:      controller,
	}
}

// TODO options

func (s *WebsocketServer) NewConnection(conn *websocket.Conn) *session.Connection {
	return s.connectionPool.Get(conn)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO
	},
}

func (s *WebsocketServer) Run(ctx context.Context) error {
	go s.stopHandler(ctx)

	s.handleConnections(ctx)

	err := http.ListenAndServe(s.address, nil)
	if err != nil {
		return fmt.Errorf("error serve http. %w", err)
	}

	return nil
}

func (s *WebsocketServer) Stop() {
	conns, free := s.connectionPool.Connections()
	defer free()

	for _, conn := range conns {
		s.log.Warn("cloging connection. ", slog.AnyValue(conn.Id()))
		if err := conn.Close(); err != nil {
			s.log.Error("error close connection %s. error: %s", conn.Id(), err.Error())
		}

		s.connectionPool.Put(conn)
	}
}

func (s *WebsocketServer) Log() *slog.Logger {
	return s.log
}

func (s *WebsocketServer) Address() string {
	return s.address
}

func freePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}

	return
}

func (s *WebsocketServer) handleConnections(ctx context.Context) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error

		wsconn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer wsconn.Close()

		conn := s.connectionPool.Get(wsconn)
		conn.AttachContext(ctx)

		conn.SetLogger(s.log.With(
			slog.String("connection", conn.Id().String()),
			slog.String("remote_addr", wsconn.RemoteAddr().String()),
		))

		go conn.Read()
	})
}

func (s *WebsocketServer) stopHandler(ctx context.Context) {
	for {
		select {
		case <-s.closeCh:
			s.log.Warn("stopped pushing workers. shutting down")
			s.Stop()

			return

		case <-ctx.Done():
			s.log.Warn("stopped pushing workers. shutting down")
			s.Stop()

			return

		case <-time.After(8 * time.Second):
			// todo helthcheck consul
		}
	}
}
