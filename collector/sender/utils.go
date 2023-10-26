package sender

import (
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/sender/azurebus"
	"github.com/krau/manyacg/collector/sender/rabbitmq"
)

func NewSender() Sender {
	switch config.Cfg.Sender.Type {
	case "azure":
		azurebus.InitAzureBus()
		return new(azurebus.SenderAzureBus)
	case "rabbitmq":
		rabbitmq.InitRabbitMQ()
		return new(rabbitmq.SenderRabbitMQ)
	}
	return nil
}
