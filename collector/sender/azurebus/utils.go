package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
)

var azureClient *azservicebus.Client

func getAzureClient() (*azservicebus.Client, error) {
	azClient, err := azservicebus.NewClientFromConnectionString(config.Cfg.Sender.Azure.BusConnectionString, nil)
	if err != nil {
		return nil, err
	}
	return azClient, err
}

func init() {
	aC, err := getAzureClient()
	if err != nil {
		logger.L.Fatalf("Error getting azure client: %s", err.Error())
		return
	}
	azureClient = aC
}
