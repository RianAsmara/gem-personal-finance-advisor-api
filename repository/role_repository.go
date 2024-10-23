package repository

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
)

type RoleRepository interface {
	GetRoles(ctx context.Context) []entity.Role
	GetRoleById(ctx context.Context, id string) (entity.Role, error)
	Insert(ctx context.Context, role entity.Role) entity.Role
	Update(ctx context.Context, role entity.Role) entity.Role
	Delete(ctx context.Context, id string)
}
