package repository

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
)

type UserRepository interface {
	GetUsers(ctx context.Context) []entity.User
	GetUserById(ctx context.Context, id string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) entity.User
	Insert(ctx context.Context, user entity.User)
	Update(ctx context.Context, user entity.User) entity.User
	Delete(ctx context.Context, id string)
	// DeleteAll()
}
