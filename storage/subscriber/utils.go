package subscriber

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/subscriber/azurebus"
	"github.com/krau/manyacg/storage/subscriber/rabbitmq"
)

func NewSubscriber() Subscriber {
	switch config.Cfg.Subscriber.Type {
	case "rabbitmq":
		rabbitmq.InitRabbitMQ()
		return new(rabbitmq.SubscriberRabbitMQ)
	case "azure":
		azurebus.InitAzureBus()
		return new(azurebus.SubscriberAzureBus)
	}
	return nil
}
