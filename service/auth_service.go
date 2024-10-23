package services

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
)

type AuthService interface {
	AuthenticationService(ctx context.Context, request model.LoginRequest) entity.User
}
