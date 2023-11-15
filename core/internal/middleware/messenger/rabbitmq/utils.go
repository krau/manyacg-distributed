package rabbitmq

import (
	"fmt"

	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitmqConn *amqp.Connection
var rabbitmqChannel *amqp.Channel
var rabbitmqSubQueue amqp.Queue
var rabbitmqDeliveries <-chan amqp.Delivery

func InitRabbitMQ() {
	if rabbitmqConn != nil && rabbitmqChannel != nil && rabbitmqSubQueue.Name != "" && rabbitmqDeliveries != nil {
		logger.L.Debug("Rabbitmq already initialized")
		return
	}
	var err error
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.Cfg.Middleware.MQ.RabbitMQ.User,
		config.Cfg.Middleware.MQ.RabbitMQ.Password,
		config.Cfg.Middleware.MQ.RabbitMQ.Host,
		config.Cfg.Middleware.MQ.RabbitMQ.Port,
		config.Cfg.Middleware.MQ.RabbitMQ.Vhost,
	)
	rabbitmqConn, err = amqp.Dial(connURL)
	if err != nil {
		logger.L.Fatalf("Error getting rabbitmq connection: %s", err.Error())
		return
	}
	rabbitmqChannel, err = rabbitmqConn.Channel()
	if err != nil {
		logger.L.Fatalf("Error getting rabbitmq channel: %s", err.Error())
		return
	}
	err = rabbitmqChannel.ExchangeDeclare(
		config.Cfg.Middleware.MQ.RabbitMQ.SubExchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error declaring rabbitmq sub exchange: %s", err.Error())
		return
	}
	err = rabbitmqChannel.ExchangeDeclare(
		config.Cfg.Middleware.MQ.RabbitMQ.PubExchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error declaring rabbitmq pub exchange: %s", err.Error())
		return
	}
	rabbitmqSubQueue, err = rabbitmqChannel.QueueDeclare(
		config.Cfg.Middleware.MQ.RabbitMQ.SubQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error declaring rabbitmq sub queue: %s", err.Error())
		return
	}
	err = rabbitmqChannel.QueueBind(
		rabbitmqSubQueue.Name,
		"",
		config.Cfg.Middleware.MQ.RabbitMQ.SubExchange,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error binding rabbitmq sub queue: %s", err.Error())
		return
	}
	autoAck := true
	if config.Cfg.App.Debug {
		autoAck = false
	}
	rabbitmqDeliveries, err = rabbitmqChannel.Consume(
		rabbitmqSubQueue.Name,
		"",
		autoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error consuming rabbitmq queue: %s", err.Error())
		return
	}
	logger.L.Info("Rabbitmq initialized")
}
