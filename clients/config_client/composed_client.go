package config_client

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/nacos_client"
)

func NewComposedConfigClient(nc nacos_client.INacosClient) (ComposedInterface, error) {
	configClient, err := NewConfigClient(nc)
	if err != nil {
		return nil, err
	}

	return struct {
		IConfigClient
		IConfigOpenApi
	}{
		IConfigOpenApi: configClient.configProxy,
		IConfigClient:  configClient,
	}, nil
}
