package main

import (
	"github.com/sculler/techtuesdayapi/database"
	"github.com/sculler/techtuesdayapi/domain"
	"github.com/sculler/techtuesdayapi/logger"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewLogger()

	pgClient, err := database.NewPostgresClient()
	if err != nil {
		log.Fatal("unable to connect to postgres database", zap.Error(err))
	}

	_ = pgClient.AutoMigrate(&domain.User{})
	_ = pgClient.AutoMigrate(&domain.TechTuesday{})

	server := CreateServer(log, pgClient)

	server.RunRouter()
}
