package services

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/model"
)

type RoleService interface {
	GetRolesService(ctx context.Context) []model.Role
	GetRoleByIdService(ctx context.Context, id string) model.Role
	CreateRoleService(ctx context.Context, request model.RoleRequest) model.RoleRequest
	UpdateRoleService(ctx context.Context, role model.RoleRequest) model.RoleRequest
	DeleteRoleService(ctx context.Context, id string)
}
