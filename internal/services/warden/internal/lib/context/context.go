package context

import (
	"context"
	"fmt"
)

type ContextKey string

var (
	CompanyIdCtxKey ContextKey = ContextKey("company_id")
	UserIdCtxKey    ContextKey = ContextKey("user_id")
)

func UserId(ctx context.Context) (string, error) {
	if userId, ok := ctx.Value(UserIdCtxKey).(string); ok {
		return userId, nil
	}

	return "", fmt.Errorf("error user id not passed in context")
}

func CompanyId(ctx context.Context) (string, error) {
	if companyId, ok := ctx.Value(CompanyIdCtxKey).(string); ok {
		return companyId, nil
	}

	return "", fmt.Errorf("error company id not passed in context")
}
