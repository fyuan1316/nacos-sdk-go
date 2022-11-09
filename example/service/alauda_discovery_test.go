package main

import (
	"fmt"
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var (
	serverIP  = "192.168.184.135"
	port      = 30421
	grpcPort  = 31254
	serverCfg []constant.ServerConfig
	clientCfg *constant.ClientConfig
	//
	nsCRM = "735c14c7-cde2-4264-8e97-3add5fab061b"
)

func setup() {
	serverCfg = []constant.ServerConfig{
		*constant.NewServerConfig(serverIP, uint64(port),
			constant.WithContextPath("/nacos"),
			constant.WithGrpcPort(uint64(grpcPort)),
		),
	}
}
func buildCC(namespace, user, password string) {
	if namespace == "" {
		namespace = "public"
	}
	clientCfg = constant.NewClientConfig(
		constant.WithNamespaceId(namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername(user),
		constant.WithPassword(password),
	)
	fmt.Printf("build clientCfg for namespace:%v, user:%v, password:%v\n", namespace, user, password)

}
func teardown() {

}
func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func Test_GetNamespaces(t *testing.T) {
	buildCC("", "nacos", "nacos")
	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  clientCfg,
			ServerConfigs: serverCfg,
		},
	)
	if err != nil {
		panic(err)
	}
	allNS, err := client.GetNamespaces()
	if err != nil {
		panic("GetNamespaces failed!")
	}
	fmt.Printf("GetNamespaces, result:%+v \n\n", allNS)
}

func Test_ListConfig(t *testing.T) {
	// as guest user
	buildCC(nsCRM, "g", "g")
	// create config client

	client, err := clients.CreateComposedConfigClient(map[string]interface{}{
		constant.KEY_CLIENT_CONFIG:  *clientCfg,
		constant.KEY_SERVER_CONFIGS: serverCfg,
	})
	if err != nil {
		panic(err)
	}
	allCFGInNS, err := client.ListConfig(nsCRM)
	if err != nil {
		panic("GetConfig failed!")
		fmt.Println(err)
	}
	fmt.Printf("GetConfig, result:%+v \n\n", allCFGInNS)
	//
	//allCFGInNS, err := client.SearchConfig(vo.SearchConfigParam{
	//	Search: "accurate",
	//})
	//if err != nil {
	//	panic("GetConfig failed!")
	//	fmt.Println(err)
	//}
	//fmt.Printf("GetConfig, result:%+v \n\n", allCFGInNS)

}
