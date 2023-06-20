package main

import (
	"os"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var nacosCli naming_client.INamingClient

func init() {
	nacosServerHost := os.Getenv("NACOS_HOST")
	if nacosServerHost != "" {
		nacosPort, err := getNacosPort()
		if err != nil {
			panic("failed to parse Nacos server port" + err.Error())
		}

		nacosCli, err = newNacosClient(nacosServerHost, nacosPort)
		if err != nil {
			panic("failed to create Nacos client" + err.Error())
		}

		svcPort, err := getSvcPort()
		if err != nil {
			panic("failed to get service port" + err.Error())
		}

		_, err = nacosCli.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          "127.0.0.1",
			Port:        svcPort,
			ServiceName: getSvcName(),
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
		})
		if err != nil {
			panic("failed to register to Nacos" + err.Error())
		}
	}
}

func getNacosPort() (uint64, error) {
	nacosServerPort := os.Getenv("NACOS_PORT")
	if nacosServerPort == "" {
		nacosServerPort = "8848"
	}

	return strconv.ParseUint(nacosServerPort, 10, 64)
}

func getSvcPort() (uint64, error) {
	svcPortStr := os.Getenv("PORT")
	if svcPortStr == "" {
		svcPortStr = "8080"
	}

	return strconv.ParseUint(svcPortStr, 10, 64)
}

func getSvcName() string {
	version := os.Getenv("VERSION")
	if version == "" {
		return "shop"
	}

	return "shop" + version
}

func newNacosClient(host string, port uint64) (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}
