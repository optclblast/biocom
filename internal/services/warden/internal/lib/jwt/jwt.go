package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/optclblast/biocom/internal/services/warden/internal/lib/models"
)

// NewToken creates new JWT token for given user
func NewToken(user models.UserIdentity, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["organization_id"] = user.OrganizationID()
	claims["exp"] = time.Now().Add(duration).Unix()

	secret := uuid.NewString() // todo proper secret using getenv or consul config value

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error sign token. %w", err)
	}

	return tokenString, nil
}
