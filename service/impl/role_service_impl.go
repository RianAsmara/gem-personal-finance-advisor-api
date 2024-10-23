package impl

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/common"
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/go-redis/redis"
)

func NewRoleServiceImpl(roleRepository *repository.RoleRepository) services.RoleService {
	return &roleServiceImpl{RoleRepository: *roleRepository}
}

type roleServiceImpl struct {
	repository.RoleRepository
	Cache *redis.Client
}

func (service *roleServiceImpl) GetRolesService(ctx context.Context) (responses []model.Role) {
	roles := service.RoleRepository.GetRoles(ctx)
	for _, role := range roles {
		responses = append(responses, model.Role{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	if len(roles) == 0 {
		return []model.Role{}
	}
	return responses
}

func (service *roleServiceImpl) GetRoleByIdService(ctx context.Context, id string) model.Role {
	roleCache := configuration.SetCache[entity.Role](service.Cache, ctx, "role", id, service.GetRoleById)

	return model.Role{
		ID:   roleCache.ID,
		Name: roleCache.Name,
	}
}

func (service *roleServiceImpl) CreateRoleService(ctx context.Context, request model.RoleRequest) model.RoleRequest {
	common.ValidateStruct(request)
	role := entity.Role{
		Name: request.Name,
	}
	service.RoleRepository.Insert(ctx, role)
	return request
}

func (service *roleServiceImpl) UpdateRoleService(ctx context.Context, request model.RoleRequest) model.RoleRequest {
	common.ValidateStruct(request)
	role := entity.Role{
		ID:   request.ID,
		Name: request.Name,
	}
	service.RoleRepository.Update(ctx, role)
	return request
}

func (service *roleServiceImpl) DeleteRoleService(ctx context.Context, id string) {
	role, err := service.RoleRepository.GetRoleById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	service.RoleRepository.Delete(ctx, role.ID)

}
