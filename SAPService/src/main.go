package main

import (
	config "SICKHackathon/SAPService/src/internal/configs"
	"SICKHackathon/SAPService/src/internal/handlers"
	"SICKHackathon/SAPService/src/internal/kafka/producer"
	"SICKHackathon/SAPService/src/internal/repositories"
	"SICKHackathon/SAPService/src/internal/services"
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

// @host localhost:8001
// @BasePath /sap-service
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

	keywordCol, err := pkg.GetMongoDbCollection(client, cfg.DBName, cfg.ColName)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	kafkaProducer := producer.NewKafkaProducer()
	//kafkaConsumerHandler := consumer.NewKafkaConsumerHandler()
	//go kafkaConsumerHandler.ConsumeKafkaTopic()

	singleClient := services.NewSingletonClient()
	SAPRepo := repositories.NewRepository(keywordCol)
	SAPService := services.NewService(SAPRepo, singleClient, kafkaProducer)
	SAPHandler := handlers.NewHandler(SAPRepo, SAPService, e)

	e.Pre(middleware.AddCorrelationID)
	e.Use(middleware.Recovery)
	e.Use(middleware.LoggingMiddleware)

	e.Validator = pkg.NewValidation()

	SAPHandler.InitEndpoints()

	log.Fatal(e.Start(cfg.BaseUrl + cfg.Port))
}
