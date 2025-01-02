package config

import (
	"fmt"
	"rest-api-go/cmd"
	handler "rest-api-go/internal/handlers"
	"rest-api-go/internal/middleware"
	"rest-api-go/internal/pkg"
	"rest-api-go/internal/repository"
	"rest-api-go/internal/router"
	"rest-api-go/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	app *gin.Engine
	log *logrus.Logger
}

func NewApp() AppConfig {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app := gin.Default()
	log := logrus.New()

	db := DBNew()

	// handle cors
	middleware.SetupCors(app)

	redisLimiter := pkg.SetupRedisLimiter("localhost:6379")

	rateLimitMiddleware := middleware.NewRateLimit(redisLimiter)
	app.Use(rateLimitMiddleware.Handler())

	// handle migrate or seeder
	cmd.MigrateOrSeed(db)

	// aws s3 to handle file
	awsS3 := pkg.NewS3AWS()

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	userService := services.NewUser(userRepository, awsS3)

	// handlers
	userHandler := handler.NewUser(log, userService)

	// Testing routes
	app.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Set up routes
	router.User(app, userHandler)
	router.TestingRouter(app)

	return AppConfig{
		app: app,
		log: log,
	}
}

func (ap *AppConfig) Run() {
	// Start the application
	if err := ap.app.Run(":8080"); err != nil {
		ap.log.Fatalf("Failed to start server: %v", err)
	}
}
