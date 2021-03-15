package service

import (
	"github.com/sculler/techtuesdayapi/domain"
	apierror "github.com/sculler/techtuesdayapi/error"
	"github.com/sculler/techtuesdayapi/logger"
	"net/http"
)

type TechTuesdayService struct {
	repository domain.ITechTuesdayRepository
	logger logger.ILogger
}

func NewTechTuesdayService(repository domain.ITechTuesdayRepository, logger logger.ILogger) TechTuesdayService {
	return TechTuesdayService{
		repository: repository,
		logger:     logger,
	}
}

func (t TechTuesdayService) Create(techTuesday *domain.TechTuesday) (int, *apierror.ApiError) {
	resp, err := t.repository.Create(techTuesday)
	if err != nil {
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (t TechTuesdayService) Delete(id int) (bool, *apierror.ApiError) {
	resp, err := t.repository.Delete(id)
	if err != nil {
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (t TechTuesdayService) GetAll() ([]domain.TechTuesday, *apierror.ApiError) {
	users, err := t.repository.GetAll()
	if err != nil {
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return users, nil
}

func (t TechTuesdayService) GetById(id int) (*domain.TechTuesday, *apierror.ApiError) {
	user, err := t.repository.GetById(id)
	if err != nil {
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return user, nil
}

func (t TechTuesdayService) Update(techTuesday *domain.TechTuesday) (bool, *apierror.ApiError) {
	resp, err := t.repository.Update(techTuesday)
	if err != nil {
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}
