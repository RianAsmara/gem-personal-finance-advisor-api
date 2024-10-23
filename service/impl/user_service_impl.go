package impl

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/common"
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/go-redis/redis"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) services.UserService {
	return &userServiceImple{UserRepository: *userRepository}
}

type userServiceImple struct {
	repository.UserRepository
	Cache *redis.Client
}

func (service *userServiceImple) GetUsersService(ctx context.Context) (responses []model.User) {
	users := service.UserRepository.GetUsers(ctx)
	for _, user := range users {
		responses = append(responses, model.User{
			ID:       user.ID,
			Email:    user.Email,
			Password: user.Password,
			// FamilyTreeID: user.FamilyTreeID.String(),
		})
	}
	if len(users) == 0 {
		return []model.User{}
	}
	return responses
}

func (service *userServiceImple) GetUserByIdService(ctx context.Context, id string) model.User {
	userCache := configuration.SetCache[entity.User](service.Cache, ctx, "user", id, service.GetUserById)

	return model.User{
		ID:       userCache.ID,
		Email:    userCache.Email,
		Password: userCache.Password,
		// FamilyTreeID: userCache.FamilyTreeID.String(),
	}
}

func (service *userServiceImple) GetUserByEmailService(ctx context.Context, email string) model.User {
	user := service.UserRepository.GetUserByEmail(ctx, email)
	return model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		// FamilyTreeID: user.FamilyTreeID.String(),
	}
}

func (service *userServiceImple) CreateUserService(ctx context.Context, model model.UserRequest) model.UserRequest {
	common.ValidateStruct(model)
	var userRoles []entity.Role
	for _, role := range model.Roles {
		userRoles = append(userRoles, entity.Role{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	user := entity.User{
		Email: model.Email,
		Roles: userRoles,
	}

	service.UserRepository.Insert(ctx, user)
	return model
}

func (service *userServiceImple) UpdateUserService(ctx context.Context, userModel model.UserRequest) model.UserRequest {
	common.ValidateStruct(userModel)
	user := entity.User{
		ID:    userModel.ID,
		Email: userModel.Email,
	}

	service.UserRepository.Update(ctx, user)

	return userModel
}

func (service *userServiceImple) DeleteUserService(ctx context.Context, id string) {
	service.UserRepository.Delete(ctx, id)
}
