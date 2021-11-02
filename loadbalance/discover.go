package loadbalance

import (
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
	return uri
}

var once sync.Once
var lb Lb

func InitDiscovery(d Discovery) {
	once.Do(func() {
		lb = NewRobinLb(d)
	})
}
