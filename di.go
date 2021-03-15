package main

import (
	"github.com/sculler/techtuesday/database"
	"github.com/sculler/techtuesday/logger"
	"github.com/sculler/techtuesday/repository"
	"github.com/sculler/techtuesday/server"
	"github.com/sculler/techtuesday/service"
)

func provideTechTuesdayRepository(pgClient *database.PostgresClient) repository.TechTuesdayRepository {
	return repository.NewTechTuesdayRepository(pgClient)
}

func provideTechTuesdayService(repository repository.TechTuesdayRepository, logger logger.ILogger) service.TechTuesdayService {
	return service.NewTechTuesdayService(repository, logger)
}

func provideUserRepository(pgClient *database.PostgresClient) repository.UserRepository {
	return repository.NewUserRepository(pgClient)
}

func provideUserService(repository repository.UserRepository, logger logger.ILogger) service.UserService {
	return service.NewUserService(repository, logger)
}

func CreateServer(logger logger.ILogger, pgClient *database.PostgresClient) server.Server {
	techTuesdayRepository := provideTechTuesdayRepository(pgClient)
	techTuesdayService := provideTechTuesdayService(techTuesdayRepository, logger)
	userRepository := provideUserRepository(pgClient)
	userService := provideUserService(userRepository, logger)

	return server.NewServer(techTuesdayService, userService, logger)
}

