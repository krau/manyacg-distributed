package azurebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

var azureClient *azservicebus.Client

func getAzureClient() (*azservicebus.Client, error) {
	azClient, err := azservicebus.NewClientFromConnectionString(config.Cfg.Subscriber.Azure.BusConnectionString, nil)
	if err != nil {
		return nil, err
	}
	return azClient, err
}

func init() {
	var err error
	azureClient, err = getAzureClient()
	if err != nil {
		logger.L.Errorf("Error getting azure client: %s", err.Error())
	}
}