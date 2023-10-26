package messenger

import (
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/messenger/azurebus"
	"github.com/krau/manyacg/core/messenger/rabbitmq"
)

func NewMessenger() Messenger {
	switch config.Cfg.Messenger.Type {
	case "azure":
		azurebus.InitAzureBus()
		return new(azurebus.MessengerAzureBus)
	case "rabbitmq":
		rabbitmq.InitRabbitMQ()
		return new(rabbitmq.MessengerRabbitMQ)
	}
	return nil
}
