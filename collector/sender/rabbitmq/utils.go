package rabbitmq

import (
	"fmt"

	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitmqConn *amqp.Connection
var rabbitmqChannel *amqp.Channel

func InitRabbitMQ() {
	if rabbitmqConn != nil && rabbitmqChannel != nil {
		logger.L.Debug("Rabbitmq already initialized")
		return
	}
	var err error
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.Cfg.Sender.RabbitMQ.User,
		config.Cfg.Sender.RabbitMQ.Password,
		config.Cfg.Sender.RabbitMQ.Host,
		config.Cfg.Sender.RabbitMQ.Port,
		config.Cfg.Sender.RabbitMQ.Vhost,
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
		config.Cfg.Sender.RabbitMQ.Exchange,
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
	logger.L.Info("Rabbitmq initialized")
}
