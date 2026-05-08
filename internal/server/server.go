package server

import (
	"auth-template/internal/config"
	"auth-template/internal/services"
	"auth-template/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg         *config.Config
	httpServer  *http.Server
	userService services.UserService
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg:         cfg,
		userService: services.NewUserService(),
	}
}

func (s *Server) Run() error {
	router := s.setupRoutes()

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.ApiPort),
		Handler: router,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) setupRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", s.HealthCheck)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", s.GetAllUsers)
			users.GET(":id", s.GetUser)
			users.POST("", s.CreateUser)
		}
	}

	return router
}

func (s *Server) HealthCheck(ctx *gin.Context) {
	utils.OK(ctx, 200, gin.H{"status": "ok"})
}
