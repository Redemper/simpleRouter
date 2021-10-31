package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"simpleRouter/loadbalance"
)

const DriverName = "nacos"

type Client struct {
	Services []*loadbalance.Service
	constant.ClientConfig
	ServerConfigs []constant.ServerConfig
	NamingClient  naming_client.INamingClient
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
		s := transferService(service)
		c.Services = append(c.Services, s)
	}
	return c.Services, nil
}
func (c *Client) Close() error {
	return nil
}
func (c *Client) Open() error {
	return nil
}

// 转换nacos的service到自定义service
func transferService(s model.Service) *loadbalance.Service {
	hosts := s.Hosts
	ins := make([]*loadbalance.Instance, len(hosts))
	for _, h := range hosts {
		var is loadbalance.InstanceStatus
		if h.Healthy {
			is = loadbalance.InstanceRun
		} else {
			is = loadbalance.InstanceDown
		}
		instance := loadbalance.NewInstance(h.InstanceId, h.Ip, h.Port, h.Weight, is)
		ins = append(ins, instance)
	}
	return loadbalance.NewService(s.Name, ins)
}
