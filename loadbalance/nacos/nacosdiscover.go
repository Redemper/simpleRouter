package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
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

func (c *Client) GetServers(mp map[string]interface{}) ([]*loadbalance.Service, error) {
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
