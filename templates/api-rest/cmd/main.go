package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.ModuleName}}/internal/config"
	"{{.ModuleName}}/internal/handler"
	"{{.ModuleName}}/internal/middleware"
	{{- if .DatabaseType}}
	"{{.ModuleName}}/internal/database"
	{{- end}}
	{{- if .CacheType}}
	"{{.ModuleName}}/internal/cache"
	{{- end}}

	"github.com/gin-gonic/gin"
	gl "github.com/rafa-mori/gocrafter/logger"
	{{- if hasFeature "API Documentation"}}
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "{{.ModuleName}}/docs"
	{{- end}}
)

// @title {{.ProjectName}} API
// @version 1.0
// @description This is the {{.ProjectName}} API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup logger
	setupLogger(cfg)

	// Setup Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	{{- if .DatabaseType}}
	// Initialize database
	db, err := database.Initialize(cfg)
	if err != nil {
		gl.Log("fatal", "Failed to initialize database: %v", err)
	}
	defer database.Close(db)
	{{- end}}

	{{- if .CacheType}}
	// Initialize cache
	cacheClient, err := cache.Initialize(cfg)
	if err != nil {
		gl.Log("fatal", "Failed to initialize cache: %v", err)
	}
	defer cache.Close(cacheClient)
	{{- end}}

	// Create dependencies
	deps := &handler.Dependencies{
		Config: cfg,
		{{- if .DatabaseType}}
		DB:     db,
		{{- end}}
		{{- if .CacheType}}
		Cache:  cacheClient,
		{{- end}}
	}

	// Setup router
	router := setupRouter(deps)

	// Create server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		gl.Log("info", "Starting server on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			gl.Log("fatal", "Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	gl.Log("info", "Shutting down server...")

	// Give the server 30 seconds to finish handling requests
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		gl.Log("fatal", "Server forced to shutdown: %v", err)
	}

	gl.Log("info", "Server stopped")
}

func setupRouter(deps *handler.Dependencies) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	{{- if hasFeature "CORS"}}
	router.Use(middleware.CORS())
	{{- end}}
	{{- if hasFeature "Rate Limiting"}}
	router.Use(middleware.RateLimit())
	{{- end}}

	// Health check endpoint
	router.GET("/health", handler.HealthCheck(deps))

	{{- if hasFeature "API Documentation"}}
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{{- end}}

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Initialize handlers
		h := handler.NewHandlers(deps)

		// Example routes
		v1.GET("/users", h.GetUsers)
		v1.POST("/users", h.CreateUser)
		v1.GET("/users/:id", h.GetUser)
		v1.PUT("/users/:id", h.UpdateUser)
		v1.DELETE("/users/:id", h.DeleteUser)
	}

	return router
}

func setupLogger(cfg *config.Config) {
	if cfg.Environment == "development" {
		gl.SetDebug(true)
	}
}
