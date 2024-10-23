package impl

import (
	"context"
	"errors"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	"gorm.io/gorm"
)

func NewRoleRepositoryImpl(DB *gorm.DB) repository.RoleRepository {
	return &roleRepositoryImpl{DB: DB}
}

type roleRepositoryImpl struct {
	*gorm.DB
}

func (repository *roleRepositoryImpl) GetRoles(ctx context.Context) []entity.Role {
	var roles []entity.Role
	repository.DB.WithContext(ctx).Find(&roles)
	return roles
}

func (repository *roleRepositoryImpl) GetRoleById(ctx context.Context, id string) (entity.Role, error) {
	var role entity.Role
	result := repository.DB.WithContext(ctx).First(&role, id)

	if result.RowsAffected == 0 {
		return entity.Role{}, errors.New("role not found")
	}
	return role, nil
}

func (repository *roleRepositoryImpl) Insert(ctx context.Context, role entity.Role) entity.Role {
	err := repository.DB.WithContext(ctx).Create(&role)
	exception.PanicLogging(err)
	return role
}

func (repository *roleRepositoryImpl) Update(ctx context.Context, role entity.Role) entity.Role {
	err := repository.DB.WithContext(ctx).Model(&role).Updates(role)
	exception.PanicLogging(err)
	return role
}

func (repository *roleRepositoryImpl) Delete(ctx context.Context, id string) {
	repository.DB.WithContext(ctx).Delete(&entity.Role{}, id)
}
