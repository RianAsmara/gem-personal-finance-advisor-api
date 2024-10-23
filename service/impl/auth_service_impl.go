package impl

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	"github.com/RianAsmara/personal-finance-advisor-api/repository"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthServiceImpl(authRepository *repository.AuthRepository) services.AuthService {
	return &authServiceImpl{AuthRepository: *authRepository}
}

type authServiceImpl struct {
	repository.AuthRepository
}

func (service *authServiceImpl) AuthenticationService(ctx context.Context, request model.LoginRequest) entity.User {
	authResult, err := service.AuthRepository.LoginRepository(ctx, request.Email)

	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(authResult.Password), []byte(request.Password))
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	return authResult
}
