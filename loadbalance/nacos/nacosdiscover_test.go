package nacos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net/http"
	"os"
	"testing"
)

var clientConfig = *constant.NewClientConfig(
	constant.WithNamespaceId("501689b2-129f-450c-8735-b04a5978b016"), //当namespace是public时，此处填空字符串。
	constant.WithTimeoutMs(5000),
	constant.WithNotLoadCacheAtStart(true),
	constant.WithLogDir("/tmp/nacos/log"),
	constant.WithCacheDir("/tmp/nacos/cache"),
	constant.WithRotateTime("1h"),
	constant.WithMaxAge(3),
	constant.WithLogLevel("debug"),
	constant.WithUsername("nacos"),
	constant.WithPassword("nacos"),
)
var serverConfigs = []constant.ServerConfig{
	{
		IpAddr:      "127.0.0.1",
		ContextPath: "/nacos",
		Port:        8848,
		Scheme:      "http",
	},
}

func TestNacos(t *testing.T) {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	info, err := namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		PageNo:   1,
		PageSize: 20,
	})
	doms := info.Doms
	for _, d := range doms {
		service, err := namingClient.GetService(vo.GetServiceParam{
			ServiceName: d,
		})
		if err != nil {
			fmt.Println("getservice error", err)
		}
		fmt.Println(service)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}

func TestRegister(t *testing.T) {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "172.16.6.149",
		Port:        6654,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		fmt.Println("register error == ", err)
		os.Exit(0)
	}
	fmt.Println(success)
	//time.Sleep(time.Minute * 10)
	router := gin.Default()
	http.ListenAndServe(":6654", router)
}
