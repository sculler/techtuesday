package service

import (
	"github.com/sculler/techtuesday/domain"
	apierror "github.com/sculler/techtuesday/error"
	"github.com/sculler/techtuesday/logger"
	"go.uber.org/zap"
	"net/http"
)

type UserService struct {
	repository domain.IUserRepository
	logger logger.ILogger
}

func NewUserService(repository domain.IUserRepository, logger logger.ILogger) UserService {
	return UserService{
		repository: repository,
		logger:     logger,
	}
}

func (u UserService) GetAll() ([]domain.User, *apierror.ApiError) {
	resp, err := u.repository.GetAll()
	if err != nil {
		u.logger.Error("unable to get all users", zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (u UserService) GetById(id int) (*domain.User, *apierror.ApiError) {
	user, err := u.repository.GetById(id)
	if err != nil {
		u.logger.Error("unable to get user by id", zap.Int("userId", id), zap.Error(err))
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}
	return user, nil
}

func (u UserService) Create(user *domain.User) (int, *apierror.ApiError) {
	resp, err := u.repository.Create(user)
	if err != nil {
		u.logger.Error("unable to create user", zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (u UserService) Update(user *domain.User) (bool, *apierror.ApiError) {
	resp, err := u.repository.Update(user)
	if err != nil {
		u.logger.Error("unable to update user", zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (u UserService) Delete(id int) (bool, *apierror.ApiError) {
	resp, err := u.repository.Delete(id)
	if err != nil {
		u.logger.Error("unable to delete user", zap.Int("userId", id), zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}
