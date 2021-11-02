package loadbalance

import (
	"strconv"
	"strings"
	"sync"
)

type InstanceStatus int

const (
	InstanceRun = iota
	InstanceDown
	InstanceChecking
)

type Discovery interface {
	SetConfig(config string) error
	Driver() string
	SetCallback(callback func(services []*Service))
	GetServers() ([]*Service, error)
	Close() error
	Open() error
}

type Service struct {
	Name      string
	instances []*Instance
	locker    sync.RWMutex
}

type Instance struct {
	InstanceID string
	IP         string
	Port       uint64
	Weight     float64
	Status     InstanceStatus
	locker     sync.RWMutex
}

func NewService(name string, instances []*Instance) *Service {
	return &Service{
		Name:      name,
		instances: instances,
		locker:    sync.RWMutex{},
	}
}

func NewInstance(instanceID string, ip string, port uint64, weight float64, status InstanceStatus) *Instance {
	return &Instance{
		InstanceID: instanceID,
		IP:         ip,
		Port:       port,
		Weight:     weight,
		Status:     status,
		locker:     sync.RWMutex{},
	}
}

func GetTagetUriByOriginUri(uri string) string {
	uri2 := strings.Replace(uri, "http", "", 1)
	split := strings.Split(uri2, "/")
	serviceName := split[0]
	targetUri := getTargetUri(serviceName)
	if len(targetUri) < 1 {
		targetUri = uri
	}
	return targetUri
}

var once sync.Once
var lb Lb

func InitDiscovery(d Discovery) {
	once.Do(func() {
		lb = NewRobinLb(d)
	})
}

func getTargetUri(serviceName string) string {
	instance := lb.GetInstance(serviceName)
	if nil != instance {
		var sb strings.Builder
		sb.WriteString(instance.IP)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(int(instance.Port)))
		return sb.String()
	}
	return ""
}
