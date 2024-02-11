package controller

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/events"
	txOutbox "github.com/optclblast/biocom/internal/services/warden/internal/infrastructure/transactional_outbox"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/jwt"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/logger"
	userrepo "github.com/optclblast/biocom/internal/services/warden/internal/usecase/repository/user"
	ssov1 "github.com/optclblast/biocom/pkg/proto/gen/warden/sso/v1"
	userv1 "github.com/optclblast/biocom/pkg/proto/gen/warden/user/v1"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	log            *slog.Logger
	tokenTTL       time.Duration
	userRepository userrepo.Repository
	txOutbox       txOutbox.TransactionalOutbox
	eventsBuilder  events.EventsBuilder
}

func NewAuthController(
	log *slog.Logger,
	tokenTTL time.Duration,
	userRepository userrepo.Repository,
	txOutbox txOutbox.TransactionalOutbox,
	eventsBuilder events.EventsBuilder,
) *AuthController {
	return &AuthController{
		log:            log,
		userRepository: userRepository,
	}
}

func (c *AuthController) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	log := c.log.With(
		slog.String("method", "sign_in"),
	)

	if err := validateSignInRequest(req); err != nil {
		log.Error("error validate request", logger.Err(err))
		return nil, fmt.Errorf("error validate request. %w", err)
	}

	user, err := c.userRepository.Get(ctx, userrepo.GetParams{
		Login: req.GetLogin(),
	})
	if err != nil {
		return nil, fmt.Errorf("error fetch user from database. %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generate password hash. %w", err)
	}

	log.Debug(
		"passwords",
		slog.Any("req_password_hash", passwordHash),
		slog.Any("stored_password_hash", user.PasswordHash),
	)

	if slices.Compare[[]byte](passwordHash, user.PasswordHash) != 0 {
		return nil, fmt.Errorf("error invalid password. %w", ErrAccessDenied)
	}

	// if ok, creake token and send it to user
	token, err := jwt.NewToken(user, c.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", logger.Err(err))
		return nil, fmt.Errorf("error create new token. %w", err)
	}

	if err = c.txOutbox.Append(ctx, c.eventsBuilder.UserSignIn(
		user.Id,
		user.OrganizationId,
		time.Now(),
	)); err != nil {
		return nil, fmt.Errorf("error add event into transactionsl outbox. %w", err)
	}

	return &ssov1.SignInResponse{
		Token: &ssov1.Token{
			Token: token,
		},
		User: &userv1.User{
			Login:     user.Login,
			Name:      user.Name,
			CreatedAt: uint64(user.CreatedAt.UnixMilli()),
			UpdatedAt: uint64(user.UpdatedAt.UnixMilli()),
			DeletedAt: uint64(user.DeletedAt.UnixMilli()),
		},
	}, nil
}

func (c *AuthController) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {
	log := c.log.With(
		slog.String("method", "sign_up"),
	)

	defer func() {
		log.Info("finished processing sign up request")
	}()

	if err := validateSignUpRequest(req); err != nil {
		return nil, fmt.Errorf("error validate request. %w", err)
	}

	log.Info(
		"processing sign up request",
		slog.String("organization_id", req.GetOrganizationId()),
	)

	return nil, nil
}

func validateSignInRequest(req *ssov1.SignInRequest) error {
	if req.GetLogin() == "" {
		return fmt.Errorf("error missing required field - login. %w", ErrMissingRequiredField)
	}

	if req.GetOrganizationId() == "" {
		return fmt.Errorf("error missing required field - organization_id. %w", ErrMissingRequiredField)
	}

	if req.GetPassword() == "" {
		return fmt.Errorf("error missing required field - password. %w", ErrMissingRequiredField)
	}

	return nil
}

func validateSignUpRequest(req *ssov1.SignUpRequest) error {
	if req.GetLogin() == "" {
		return fmt.Errorf("error missing required field - login. %w", ErrMissingRequiredField)
	}

	if req.GetPassword() == "" {
		return fmt.Errorf("error missing required field - password. %w", ErrMissingRequiredField)
	}

	if req.GetOrganizationId() == "" {
		return fmt.Errorf("error missing required field - organization_id. %w", ErrMissingRequiredField)
	}

	return nil
}
