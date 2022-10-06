package main

import (
	config "SICKHackathon/chatbotService/src/internal/configs"
	"SICKHackathon/chatbotService/src/internal/handlers"
	"SICKHackathon/chatbotService/src/internal/repositories"
	"SICKHackathon/chatbotService/src/internal/services"
	"github.com/erenkaratas99/COApiCore/pkg"
	"github.com/erenkaratas99/COApiCore/pkg/middleware"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	pkg.InitLogrusConfig()
}

// @title Customer API
// @version 1.0
// @description This is an order server to handle the requests about CRUD ops and Order Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Eren Karataş
// @contact.url linkedin.com/in/erenkaratass
// @contact.email karatas18@itu.edu.tr

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /customer
// @schemes http
func main() {
	real_env := "dev"
	env := os.Getenv("RUN_ENVIRONMENT")
	if env == "docker" {
		real_env = "prod"
	}
	cfg, err := config.GetConfig(real_env)
	if err != nil {
		log.Fatal(err)
	}
	client, err := pkg.GetMongoClient(cfg.MongoDuration, cfg.MongoClientURI)
	if err != nil {
		log.Fatal(err)
	}

	msgCol, err := pkg.GetMongoDbCollection(client, cfg.DBName, cfg.ColName)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	singleClient := services.NewSingletonClient()
	orderRepo := repositories.NewRepository(msgCol)
	orderService := services.NewService(orderRepo, singleClient)
	orderHandler := handlers.NewHandler(orderRepo, orderService, e)

	e.Pre(middleware.AddCorrelationID)
	e.Use(middleware.Recovery)
	e.Use(middleware.LoggingMiddleware)

	e.Validator = pkg.NewValidation()

	orderHandler.InitEndpoints()

	log.Fatal(e.Start(cfg.BaseUrl + cfg.Port))
}
