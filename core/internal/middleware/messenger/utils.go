package messenger

import (
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/middleware/messenger/azurebus"
	"github.com/krau/manyacg/core/internal/middleware/messenger/rabbitmq"
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
