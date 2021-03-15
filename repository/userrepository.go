package repository

import (
	"github.com/sculler/techtuesdayapi/database"
	"github.com/sculler/techtuesdayapi/domain"
)

type UserRepository struct {
	pgClient *database.PostgresClient
}

func NewUserRepository(pgClient *database.PostgresClient) UserRepository {
	return UserRepository {
		pgClient: pgClient,
	}
}

func (u UserRepository) Create(user *domain.User) (int, error) {
	result := u.pgClient.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

func (u UserRepository) Delete(id int) (bool, error) {
	result := u.pgClient.Delete(&domain.User{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := u.pgClient.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u UserRepository) GetById(id int) (*domain.User, error) {
	var user *domain.User
	result := u.pgClient.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u UserRepository) Update(user *domain.User) (bool, error) {
	result := u.pgClient.Save(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
