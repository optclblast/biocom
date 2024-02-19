package session

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/optclblast/biocom/internal/services/ws_server/internal/usecase/controllers"
	"github.com/optclblast/biocom/internal/services/ws_server/pkg/models"
	"github.com/optclblast/biocom/internal/services/ws_server/pkg/reqctx"
	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

type SessionPool interface {
	Get(sid string) Session
	Put(ses Session)
}

type sessionPool struct {
	pool           sync.Pool
	sessionStorage SessionStorage
}

func NewSessionPool() SessionPool {
	return &sessionPool{
		pool: sync.Pool{
			New: func() any {
				return &session{}
			},
		},
	}
}

func (s *sessionPool) Get(sid string) Session {
	if sid != "" {
		ses, err := s.sessionStorage.Get(context.Background(), sid)
		if err == nil {
			return ses
		}
	}

	return s.pool.Get().(*session)
}

func (s *sessionPool) Put(ses Session) {
	s.pool.Put(ses)
}

type Session interface {
	Id() string
	Pts() uint64
	CompanyId() string
	User() *models.User
	Value(key any) any
	Deadline() time.Time
	SetUser(user *models.User)
	SetCompanyId(companyId string)
	Write(resp *apiv1.Response) error
	DispatchRequest(req reqctx.RequestContext) error
	ToProto() []byte
}

type session struct {
	mu         sync.Mutex
	authorized bool
	id         string
	companyId  string
	user       *models.User
	ctx        context.Context
	reqctx     reqctx.RequestContext
	ctxPool    reqctx.RequestContextPool
	deadline   time.Time

	pts  atomic.Uint64
	conn *Connection // nocopy

	controller *controllers.RootController
}

func (s *session) Id() string {
	return s.id
}

func (s *session) User() *models.User {
	return s.user
}

func (s *session) CompanyId() string {
	return s.companyId
}

func (s *session) Value(key any) any {
	return s.ctx.Value(key)
}

func (s *session) Deadline() time.Time {
	return s.deadline
}

func (s *session) Pts() uint64 {
	return s.pts.Load()
}

func (s *session) SetUser(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.user = user
}

func (s *session) SetCompanyId(companyId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.companyId = companyId
}

func (s *session) Write(resp *apiv1.Response) error {
	return s.conn.Write(resp)
}

func (s *session) DispatchRequest(req reqctx.RequestContext) error {
	switch {
	case req.Request().GetAuthSignIn() != nil:
		return s.handle(req, s.controller.AuthController.SignIn, "sign_in")
	case req.Request().GetAuthSignUp() != nil:
		return s.handle(req, s.controller.AuthController.SignUp, "sign_up")
	default:
		return nil // todo return error
	}
}

func (s *session) handle(
	req reqctx.RequestContext,
	handler func(req reqctx.RequestContext) (*apiv1.Response, error),
	handlerName string,
	middlewares ...func(req reqctx.RequestContext) error,
) error {
	for _, mw := range middlewares {
		if err := mw(req); err != nil {
			return err
		}
	}

	resp, err := handler(req)
	if err != nil {
		return fmt.Errorf("error handle request. %w", err)
	}

	return s.Write(resp)
}

func (s *session) authMiddleware() func(req reqctx.RequestContext) error {
	return func(req reqctx.RequestContext) error {
		if !s.authorized {
			return fmt.Errorf("error unauthorized") // todo wrap typed error
		}

		return nil
	}
}

// todo destory session

func FromProto(sesProto *apiv1.Session) Session {
	return &session{
		id: sesProto.GetId(),
		//userId:    sesProto.GetUserId(), // set user, mapped from proto warden model
		companyId: sesProto.GetCompanyId(),
		// todo map values
	}
}

func (s *session) ToProto() []byte {
	return []byte{} // TODO
}

func (s *session) newRequestContext(
	req *apiv1.Request,
	ctx context.Context,
) reqctx.RequestContext {
	return s.ctxPool.Get(req, s.pts.Load(), ctx)
}

func (s *session) freeRequestContext() {
	s.ctxPool.Put(s.reqctx)
}
