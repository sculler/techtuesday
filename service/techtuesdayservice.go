package service

import (
	"github.com/sculler/techtuesday/domain"
	apierror "github.com/sculler/techtuesday/error"
	"github.com/sculler/techtuesday/logger"
	"go.uber.org/zap"
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

func (t TechTuesdayService) GetAll() ([]domain.TechTuesday, *apierror.ApiError) {
	techTuesdays, err := t.repository.GetAll()
	if err != nil {
		t.logger.Error("unable to get all tech tuesdays", zap.Error(err))
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return techTuesdays, nil
}

func (t TechTuesdayService) GetById(id int) (*domain.TechTuesday, *apierror.ApiError) {
	techTuesday, err := t.repository.GetById(id)
	if err != nil {
		t.logger.Error("unable to get tech tuesday by id", zap.Int("techTuesdayId", id), zap.Error(err))
		return nil, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return techTuesday, nil
}

func (t TechTuesdayService) Create(techTuesday *domain.TechTuesday) (int, *apierror.ApiError) {
	resp, err := t.repository.Create(techTuesday)
	if err != nil {
		t.logger.Error("unable to create tech tuesday", zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}

func (t TechTuesdayService) Update(techTuesday *domain.TechTuesday) (bool, *apierror.ApiError) {
	resp, err := t.repository.Update(techTuesday)
	if err != nil {
		t.logger.Error("unable to update tech tuesday", zap.Error(err))
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
		t.logger.Error("unable to delete tech tuesday", zap.Int("techTuesdayId", id), zap.Error(err))
		return resp, &apierror.ApiError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return resp, nil
}
