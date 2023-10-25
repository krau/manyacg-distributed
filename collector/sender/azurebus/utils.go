package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
)

var azureClient *azservicebus.Client
var azureSender *azservicebus.Sender

func init() {
	var err error
	azureClient, err = azservicebus.NewClientFromConnectionString(config.Cfg.Sender.Azure.BusConnectionString, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure client: %s", err.Error())
		return
	}
	azureSender, err = azureClient.NewSender(config.Cfg.Sender.Azure.Topic, nil)
	if err != nil {
		logger.L.Fatalf("Error getting azure sender: %s", err.Error())
		return
	}
}
