package repository

import (
	"github.com/sculler/techtuesday/database"
	"github.com/sculler/techtuesday/domain"
)

type TechTuesdayRepository struct {
	pgClient *database.PostgresClient
}

func NewTechTuesdayRepository(pgClient *database.PostgresClient) TechTuesdayRepository {
	return TechTuesdayRepository {
		pgClient: pgClient,
	}
}

func (t TechTuesdayRepository) GetAll() ([]domain.TechTuesday, error){
	var techTuesdays []domain.TechTuesday
	result := t.pgClient.Preload("User").Find(&techTuesdays)
	if result.Error != nil {
		return nil, result.Error
	}
	return techTuesdays, nil
}

func (t TechTuesdayRepository) GetById(id int) (*domain.TechTuesday, error) {
	var techTuesday *domain.TechTuesday
	result := t.pgClient.Preload("User").First(&techTuesday, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return techTuesday, nil
}

func (t TechTuesdayRepository) Create(techTuesday *domain.TechTuesday) (int, error) {
	result := t.pgClient.Create(techTuesday)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(techTuesday.ID), nil
}

func (t TechTuesdayRepository) Update(techTuesday *domain.TechTuesday) (bool, error) {
	result := t.pgClient.Save(techTuesday)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (t TechTuesdayRepository) Delete(id int) (bool, error) {
	result := t.pgClient.Delete(&domain.TechTuesday{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}