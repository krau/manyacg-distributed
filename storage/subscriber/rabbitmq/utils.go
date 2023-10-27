package rabbitmq

import (
	"fmt"

	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
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
		config.Cfg.Subscriber.RabbitMQ.User,
		config.Cfg.Subscriber.RabbitMQ.Password,
		config.Cfg.Subscriber.RabbitMQ.Host,
		config.Cfg.Subscriber.RabbitMQ.Port,
		config.Cfg.Subscriber.RabbitMQ.Vhost,
	)
	rabbitmqConn, err = amqp.Dial(connURL)
	if err != nil {
		logger.L.Fatalf("Error connecting to rabbitmq: %s", err.Error())
		return
	}
	rabbitmqChannel, err = rabbitmqConn.Channel()
	if err != nil {
		logger.L.Fatalf("Error getting rabbitmq channel: %s", err.Error())
		return
	}
	rabbitmqChannel.Qos(
		30,
		0,
		false,
	)
	err = rabbitmqChannel.ExchangeDeclare(
		config.Cfg.Subscriber.RabbitMQ.SubExchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error declaring rabbitmq exchange: %s", err.Error())
		return
	}
	rabbitmqSubQueue, err = rabbitmqChannel.QueueDeclare(
		config.Cfg.Subscriber.RabbitMQ.SubQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error declaring rabbitmq queue: %s", err.Error())
		return
	}
	err = rabbitmqChannel.QueueBind(
		rabbitmqSubQueue.Name,
		"",
		config.Cfg.Subscriber.RabbitMQ.SubExchange,
		false,
		nil,
	)
	if err != nil {
		logger.L.Fatalf("Error binding rabbitmq queue: %s", err.Error())
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
