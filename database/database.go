package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	*gorm.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	dsn := "host=localhost user=tt_user password=iUj398ajesj3K dbname=tech_tuesday"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresClient{
		DB: db,
	}, nil
}