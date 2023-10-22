package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

var azureClient *azservicebus.Client
var azSubscriber *azservicebus.Receiver

func init() {
	var err error
	azureClient, err = azservicebus.NewClientFromConnectionString(config.Cfg.Subscriber.Azure.BusConnectionString, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure client: %s", err.Error())
		return
	}
	azSubscriber, err = azureClient.NewReceiverForSubscription(config.Cfg.Subscriber.Azure.SubTopic, config.Cfg.Subscriber.Azure.Subscription, nil)
	if err != nil {
		logger.L.Errorf("Error getting azure receiver: %s", err.Error())
		return
	}
}
