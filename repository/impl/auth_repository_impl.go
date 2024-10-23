package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	"gorm.io/gorm"
)

func NewAuthRepositoryImpl(DB *gorm.DB) repository.AuthRepository {
	return &authRepositoryImpl{DB: DB}
}

type authRepositoryImpl struct {
	*gorm.DB
}

func (repository *authRepositoryImpl) LoginRepository(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User

	fmt.Println("===============================")
	fmt.Println("email", email)

	result := repository.DB.WithContext(ctx).
		Where("users.email = ?", email).
		Preload("Roles").
		Find(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (repository *authRepositoryImpl) GetUserById(ctx context.Context, id string) (entity.User, error) {
	var userResult entity.User

	result := repository.DB.WithContext(ctx).
		Joins("LEFT JOIN roles ON users.role_id = roles.id").
		Where("users.id = ?", id).
		Select("users.*, roles.name as name").
		First(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}
