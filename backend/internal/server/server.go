package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/freddyamarante/go-svelte-finance-tracker/internal/config"
	"github.com/freddyamarante/go-svelte-finance-tracker/internal/handlers"
)

// Server represents the HTTP server
type Server struct {
	router *gin.Engine
	config *config.Config
}

// New creates a new server instance
func New(cfg *config.Config) *Server {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if cfg.Logging.Requests {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	// Setup CORS
	corsConfig := cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     cfg.CORS.AllowMethods,
		AllowHeaders:     cfg.CORS.AllowHeaders,
		ExposeHeaders:    cfg.CORS.ExposeHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           time.Duration(cfg.CORS.MaxAge) * time.Second,
	}

	if cfg.Logging.CORS {
		log.Printf("CORS Configuration: %+v", corsConfig)
	}

	router.Use(cors.New(corsConfig))

	// Setup routes
	setupRoutes(router)

	return &Server{
		router: router,
		config: cfg,
	}
}

// setupRoutes configures all the routes
func setupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Transaction routes
	router.GET("/transactions", handlers.GetTransactions)
	router.POST("/transactions", handlers.CreateTransaction)
}

// Start starts the server
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
	log.Printf("Server starting on %s", addr)
	return s.router.Run(addr)
} 