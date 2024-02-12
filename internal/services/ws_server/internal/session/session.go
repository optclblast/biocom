package session

import (
	"sync"

	apiv1 "github.com/optclblast/biocom/pkg/proto/gen/ws/api"
)

type Session interface {
}

type session struct {
	mu        sync.Mutex
	id        string
	userId    string
	companyId string
	values    map[string]any
	valuesMu  sync.RWMutex
}

func FromProto(sesProto *apiv1.Session) Session {
	return &session{
		id:        sesProto.GetId(),
		userId:    sesProto.GetUserId(),
		companyId: sesProto.GetCompanyId(),
		// todo map values
	}
}
