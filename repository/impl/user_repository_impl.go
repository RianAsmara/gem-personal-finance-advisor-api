package impl

import (
	"context"
	"errors"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (repository *userRepositoryImpl) GetUsers(ctx context.Context) []entity.User {
	var users []entity.User
	repository.DB.WithContext(ctx).Find(&users)
	return users
}

func (repository *userRepositoryImpl) GetUserById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	result := repository.DB.WithContext(ctx).First(&user, id)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}

func (repository *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) entity.User {
	var user entity.User
	repository.DB.WithContext(ctx).Where("email = ?", email).First(&user)
	return user
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) {
	var roles []entity.Role

	for _, role := range user.Roles {
		roles = append(roles, entity.Role{
			ID:   uuid.New().String(),
			Name: role.Name,
		})
	}

	user = entity.User{
		ID:       uuid.New().String(),
		Email:    user.Email,
		Password: user.Password,
		IsActive: user.IsActive,
		Roles:    roles,
	}

	err := repository.DB.Create(&user).Error
	exception.PanicLogging(err)

}

func (repository *userRepositoryImpl) Update(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Model(&user).Updates(user).Error
	exception.PanicLogging(err)
	return user
}

func (repository *userRepositoryImpl) Delete(ctx context.Context, id string) {
	err := repository.DB.WithContext(ctx).Delete(&entity.User{}, id).Error
	exception.PanicLogging(err)
}

// func (repository *userRepositoryImpl) DeleteAll() {
// 	err := repository.DB.WithContext(ctx).Where("1 = 1").Delete(&entity.User{}).Error
// 	exception.PanicLogging(err)
// }
