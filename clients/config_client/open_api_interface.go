package config_client

import "github.com/nacos-group/nacos-sdk-go/v2/model"

type IConfigOpenApi interface {
	ListConfig(tenant string) (*model.ConfigPage, error)
}
