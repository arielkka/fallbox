package app

import (
	"fmt"
	"os"

	"github.com/arielkka/fallbox/handler/internal/broker"
	"github.com/arielkka/fallbox/handler/internal/handler"
	"github.com/arielkka/fallbox/handler/internal/service"
	"github.com/arielkka/fallbox/handler/pkg/mysql"

	"github.com/arielkka/fallbox/handler/config"
	logger "github.com/arielkka/fallbox/handler/pkg/logrus"
	"github.com/arielkka/rabbitmq"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./handler.env"); err != nil {
		fmt.Println("No .env file found", err)
	}
}

func Run() error {
	log := logger.NewLogrus(os.Getenv("LOGGER_PATH_HANDLER") + "/handler.log")
	log.Info("logger was initialized")

	serviceConfig, err := config.NewServiceConfig(os.Getenv("CONFIG_PATH_HANDLER"))
	if err != nil {
		log.Fatalf("Couldn't initialize service; error = %v", err)
	}
	log.Info("service config was initialized")

	mqConfig, err := config.NewRabbitMQConfig(os.Getenv("CONFIG_PATH_HANDLER"))
	if err != nil {
		log.Fatalf("Couldn't initialize service; error = %v", err)
	}
	log.Info("broker config was initialized")

	rabbitmqClient, err := rabbitmq.NewClient(serviceConfig.Service.Name, mqConfig)
	if err != nil {
		log.Fatalf("Couldn't create rabbitmq client; error = %v", err)
	}
	log.Info("rabbitmq client was initialized")

	sqlConnection, err := mysql.NewMySQL(serviceConfig)
	if err != nil {
		log.Fatalf("Couldn't connect to database; error = %v", err.Error())
	}
	log.Info("sql connection was initialized")

	rabbitMQ := broker.NewBroker(rabbitmqClient, serviceConfig)

	err = rabbitMQ.CreateExchanges()
	if err != nil {
		log.Fatalf("Couldn't create exchanges; error: %v", err.Error())
	}
	log.Info("created exchanges")

	err = rabbitMQ.CreateQueues()
	if err != nil {
		log.Fatalf("Couldn't create queues; error: %v", err.Error())
	}
	log.Info("created queues")

	err = rabbitMQ.CreateListeners()
	if err != nil {
		log.Fatalf("Couldn't create listeners; error: %v", err.Error())
	}
	log.Info("created listeners")

	storage := service.NewStorage(sqlConnection)
	log.Info("created storage")

	serv := service.NewService(serviceConfig, rabbitMQ, storage)
	log.Info("created service")

	return handler.NewRouter(serviceConfig, serv).Run()
}
