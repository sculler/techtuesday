package service

import (
	"github.com/sculler/techtuesdayapi/domain"
	apierror "github.com/sculler/techtuesdayapi/error"
	"github.com/sculler/techtuesdayapi/logger"
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

func (u UserService) Create(user *domain.User) (int, *apierror.ApiError) {
	resp, err := u.repository.Create(user)
	if err != nil {
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
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (u UserService) GetAll() ([]domain.User, *apierror.ApiError) {
	resp, err := u.repository.GetAll()
	if err != nil {
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
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}
	return user, nil
}

func (u UserService) Update(user *domain.User) (bool, *apierror.ApiError) {
	resp, err := u.repository.Update(user)
	if err != nil {
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

