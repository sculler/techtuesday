package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sculler/techtuesday/domain"
	"github.com/sculler/techtuesday/handler"
	"github.com/sculler/techtuesday/logger"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	techTuesdayService domain.ITechTuesdayService
	userService domain.IUserService
	logger logger.ILogger
}

func NewServer(
	techTuesdayService domain.ITechTuesdayService,
	userService domain.IUserService,
	logger logger.ILogger,
) Server {
	router := gin.New()

	server := Server{
		router:             router,
		techTuesdayService: techTuesdayService,
		userService:        userService,
		logger:             logger,
	}

	server.setupRoutes()

	return server
}

func (s Server) setupRoutes() {

	// Initialize Handlers
	techTuesdayHandler := handler.TechTuesdayHandler{
		TechTuesdayService: s.techTuesdayService,
	}

	userHandler := handler.UserHandler{
		UserService: s.userService,
	}

	// TechTuesday endpoints
	s.router.GET("techtuesday", techTuesdayHandler.HandleTechTuesdayGetAll())
	s.router.GET("techtuesday/:techTuesdayId", techTuesdayHandler.HandleTechTuesdayGetById())
	s.router.POST("techtuesday", techTuesdayHandler.HandleTechTuesdayCreate())
	s.router.PUT("techtuesday/:techTuesdayId", techTuesdayHandler.HandleTechTuesdayUpdate())
	s.router.DELETE("techtuesday/:techTuesdayId", techTuesdayHandler.HandleTechTuesdayDelete())

	// User endpoints
	s.router.GET("user", userHandler.HandlerUserGetAll())
	s.router.GET("user/:userId", userHandler.HandleUserGetById())
	s.router.POST("user", userHandler.HandleUserCreate())
	s.router.PUT("user/:userId", userHandler.HandleUserUpdate())
	s.router.DELETE("user/:userId", userHandler.HandleUserDelete())
}

func (s Server) RunRouter() {
	err := s.router.Run()
	if err != nil {
		s.logger.Fatal("could not start routing engine", zap.Error(err))
	}
}