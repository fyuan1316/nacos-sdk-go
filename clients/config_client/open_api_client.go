package config_client

import (
	"errors"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/common/nacos_server"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var _ IConfigOpenApi = &ConfigOpenApiClient{}

type ConfigOpenApiClient struct {
	NacosServer  *nacos_server.NacosServer
	ClientConfig *constant.ClientConfig
	ConfigClient *ConfigClient
}

func (oac *ConfigOpenApiClient) ListConfig(tenant string) (*model.ConfigPage, error) {
	if tenant == "" {
		return nil, errors.New("[client.ListConfig] param.tenant must not be empty")
	}
	return oac.ConfigClient.searchConfigInnerByNamespaceID(tenant, vo.SearchConfigParam{
		Search: "accurate",
	})
}
