package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/logger"
)

var azureClient *azservicebus.Client
var azureSender *azservicebus.Sender
var azureSubscriber *azservicebus.Receiver

func init() {
	var err error
	azureClient, err = azservicebus.NewClientFromConnectionString(config.Cfg.Messenger.Azure.BusConnectionString, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure client: %s", err.Error())
		return
	}
	azureSender, err = azureClient.NewSender(config.Cfg.Messenger.Azure.PubTopic, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure sender: %s", err.Error())
		return
	}
	azureSubscriber, err = azureClient.NewReceiverForSubscription(config.Cfg.Messenger.Azure.SubTopic, config.Cfg.Messenger.Azure.Subscription, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure receiver: %s", err.Error())
		return
	}
}
