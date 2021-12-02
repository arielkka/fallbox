package app

import (
	"fmt"
	"github.com/arielkka/fallbox/excel/config"
	"github.com/arielkka/fallbox/excel/internal/broker"
	"github.com/arielkka/fallbox/excel/internal/service"
	logger "github.com/arielkka/fallbox/excel/pkg/logrus"
	"github.com/arielkka/fallbox/excel/pkg/mysql"
	"github.com/arielkka/rabbitmq"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load("./handler.env"); err != nil {
		fmt.Println("No .env file found", err)
	}
}

func Run() {
	log := logger.NewLogrus(os.Getenv("LOGGER_PATH_EXCEL"))
	log.Println("logger initialized")

	cfg, err := config.NewServiceConfig(os.Getenv("CONFIG_PATH_EXCEL"))
	if err != nil {
		log.Fatalf("Couldn't create excel config; error = %v", err)
	}

	brokerConfig, err := config.NewRabbitMQConfig(os.Getenv("CONFIG_PATH_EXCEL"))
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
