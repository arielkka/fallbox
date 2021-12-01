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
	log.Info("logger was initialized", " ", os.Getenv("LOGGER_PATH_HANDLER"), " ", os.Getenv("CONFIG_PATH_HANDLER"))

	serviceConfig, err := config.NewServiceConfig(os.Getenv("CONFIG_PATH_HANDLER"))
	if err != nil {
		log.Fatalf("Couldn't initialize service; error = %v", err)
	}

	mqConfig, err := config.NewRabbitMQConfig(os.Getenv("CONFIG_PATH_HANDLER"))
	if err != nil {
		log.Fatalf("Couldn't initialize service; error = %v", err)
	}

	rabbitmqClient, err := rabbitmq.NewClient(serviceConfig.Service.Name, mqConfig)
	if err != nil {
		log.Fatalf("Couldn't create rabbitmq client; error = %v", err)
	}

	sqlConnection, err := mysql.NewMySQL(serviceConfig)
	if err != nil {
		log.Fatalf("Couldn't connect to database; error = %v", err.Error())
	}

	rabbitMQ := broker.NewBroker(rabbitmqClient, serviceConfig)

	storage := service.NewStorage(sqlConnection)

	serv := service.NewService(serviceConfig, rabbitMQ, storage)

	return handler.NewRouter(serviceConfig, serv).Run()
}
