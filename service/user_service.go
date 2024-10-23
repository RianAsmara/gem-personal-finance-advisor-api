package services

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/model"
)

type UserService interface {
	GetUsersService(ctx context.Context) []model.User
	GetUserByIdService(ctx context.Context, id string) model.User
	GetUserByEmailService(ctx context.Context, email string) model.User
	CreateUserService(ctx context.Context, request model.UserRequest) model.UserRequest
	UpdateUserService(ctx context.Context, user model.UserRequest) model.UserRequest
	DeleteUserService(ctx context.Context, id string)
}
