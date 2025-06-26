package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Info("No .env file found")
	}

	// Setup logging
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "{{port}}"
	}

	// Setup Gin router
	r := gin.Default()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Routes
	setupRoutes(r)

	// Start server
	logrus.Infof("Starting {{project_name}} server on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "{{project_name}}",
			"version": "1.0.0",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			name := c.DefaultQuery("name", "World")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Hello, %s!", name),
			})
		})

		v1.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"project":     "{{project_name}}",
				"description": "{{description}}",
				"author":      "{{author}}",
				"license":     "{{license}}",
				"go_version":  "{{go_version}}",
			})
		})
	}
}
