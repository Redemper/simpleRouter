package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"simpleRouter/loadbalance"
)

const DriverName = "nacos"

type Client struct {
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
	//nacosConf, err := conf.ReadNacosConfFromYaml()
	//if err != nil {
	//	log.Fatal("cant load nacos conf ,err == ",err)
	//}
	//cc := nacosConf.Cc
	//sc := nacosConf.Sc
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
		log.Println(service)
	}
	if err != nil {
		log.Println(err)
	}
	log.Println(info)

	//ns := mp["namespace"].(string)
	//group := mp["group"].(string)
	//servicesInfo, err := c.NamingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
	//    PageNo:    1,
	//    PageSize:  20,
	//    NameSpace: ns,
	//    GroupName: group,
	//})
	//if err != nil {
	//    return nil, err
	//}
	return nil, nil
	//
	//instances, err := c.NamingClient.SelectInstances(vo.SelectInstancesParam{
	//    ServiceName: "demo.go",
	//    GroupName:   "group-a",             // 默认值DEFAULT_GROUP
	//    Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	//})
	//if err != nil {
	//    return nil, err
	//}
	//for ins := range instances {
	//
	//}
}
func (c *Client) Close() error {
	return nil
}
func (c *Client) Open() error {
	return nil
}
