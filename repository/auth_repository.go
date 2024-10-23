package repository

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
)

type AuthRepository interface {
	LoginRepository(ctx context.Context, email string) (entity.User, error)
}
