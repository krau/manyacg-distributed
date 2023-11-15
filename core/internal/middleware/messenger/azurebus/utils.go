package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
)

var azureClient *azservicebus.Client
var azureSender *azservicebus.Sender
var azureSubscriber *azservicebus.Receiver

func InitAzureBus() {
	if azureClient != nil && azureSender != nil && azureSubscriber != nil {
		logger.L.Debug("Azure bus already initialized")
		return
	}
	var err error
	azureClient, err = azservicebus.NewClientFromConnectionString(config.Cfg.Middleware.MQ.Azure.BusConnectionString, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure client: %s", err.Error())
		return
	}
	azureSender, err = azureClient.NewSender(config.Cfg.Middleware.MQ.Azure.PubTopic, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure sender: %s", err.Error())
		return
	}
	azureSubscriber, err = azureClient.NewReceiverForSubscription(config.Cfg.Middleware.MQ.Azure.SubTopic, config.Cfg.Middleware.MQ.Azure.Subscription, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure receiver: %s", err.Error())
		return
	}
	logger.L.Info("Azure bus initialized")
}
