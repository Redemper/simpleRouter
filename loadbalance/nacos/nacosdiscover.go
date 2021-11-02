package nacos

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"simpleRouter/loadbalance"
	"sync"
)

const DriverName = "nacos"

type Client struct {
	Services     []*loadbalance.Service
	NamingClient naming_client.INamingClient
	once         sync.Once
}

func (c *Client) SetConfig(config string) error {
	return nil
}

func (c *Client) Driver() string {
	return DriverName
}

func (c *Client) SetCallback(callback func(services []*loadbalance.Service)) {
	return
}

func (c *Client) GetServers() ([]*loadbalance.Service, error) {
	if nil == c.NamingClient {
		return nil, errors.New("nacos client is not ready")
	}
	c.once.Do(func() {
		info, err := c.NamingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
			PageNo:   1,
			PageSize: 20,
		})
		if err != nil || &info == nil {
			log.Fatal("cant get service info")
		}
		doms := info.Doms
		for _, d := range doms {
			service, err := c.NamingClient.GetService(vo.GetServiceParam{
				ServiceName: d,
			})
			if err != nil {
				log.Println("getservice error", err)
			}
			s := transferService(d, service)
			c.Services = append(c.Services, s)
		}
	})
	return c.Services, nil
}
func (c *Client) Close() error {
	return nil
}
func (c *Client) Open() error {
	return nil
}

// 转换nacos的service到自定义service
func transferService(dom string, s model.Service) *loadbalance.Service {
	hosts := s.Hosts
	ins := make([]*loadbalance.Instance, len(hosts))
	for i, h := range hosts {
		var is loadbalance.InstanceStatus
		if h.Healthy {
			is = loadbalance.InstanceRun
		} else {
			is = loadbalance.InstanceDown
		}
		instance := loadbalance.NewInstance(h.InstanceId, h.Ip, h.Port, h.Weight, is)
		ins[i] = instance
	}
	return loadbalance.NewService(dom, ins)
}

func init() {
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
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	c := &Client{
		NamingClient: namingClient,
	}
	loadbalance.InitDiscovery(c)

}
