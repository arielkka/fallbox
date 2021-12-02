package app

import (
	"fmt"
	"os"

	"github.com/arielkka/fallbox/txt/config"
	"github.com/arielkka/fallbox/txt/internal/broker"
	"github.com/arielkka/fallbox/txt/internal/service"
	logger "github.com/arielkka/fallbox/txt/pkg/logrus"
	"github.com/arielkka/fallbox/txt/pkg/mysql"
	"github.com/arielkka/fallbox/txt/pkg/rabbitmq"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./txt.env"); err != nil {
		fmt.Println("No .env file found", err)
	}
}

func Run() {
	log := logger.NewLogrus(os.Getenv("LOGGER_PATH_TXT") + "/txt.log")
	log.Println("logger initialized")

	cfg, err := config.NewServiceConfig(os.Getenv("CONFIG_PATH_TXT"))
	if err != nil {
		log.Fatalf("Couldn't create excel config; error = %v", err)
	}

	brokerConfig, err := config.NewRabbitMQConfig(os.Getenv("CONFIG_PATH_TXT"))
	if err != nil {
		log.Fatalf("Couldn't create rabbitmqClient config; error = %v", err)
	}

	rabbitmqClient, err := rabbitmq.NewClient(cfg.Service.Name, brokerConfig)
	if err != nil {
		log.Fatalf("Couldn't create rabbitmqClient client; error = %v", err)
	}

	rabbitMQ := broker.NewBroker(rabbitmqClient, cfg)
	sqlConnect, err := mysql.NewMySQL(cfg)
	if err != nil {
		log.Fatalf("Couldn't connect to database; error = %v", err.Error())
	}
	store := service.NewStorage(sqlConnect)
	serv := service.NewService(cfg, rabbitMQ, store, log)
	err = rabbitMQ.CreateQueues()
	if err != nil {
		log.Fatalf("Couldn't create queues; error: %v", err.Error())
	}
	err = rabbitMQ.CreateListeners()
	if err != nil {
		log.Fatalf("Couldn't create listeners; error: %v", err.Error())
	}
	forever := make(chan struct{})
	go serv.Run()
	<-forever
}
