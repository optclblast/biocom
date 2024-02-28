package controller

import (
	"log/slog"

	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/events"
	txOutbox "github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/transactional_outbox"
	"github.com/optclblast/biocom/internal/services/warden/internal/usecase/repository/permissions"
)

type PermissionsController struct {
	log             *slog.Logger
	txOutbox        txOutbox.TransactionalOutbox
	eventsBuilder   events.EventsBuilder
	permissionsRepo permissions.Repository
}
