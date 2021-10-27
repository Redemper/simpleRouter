package loadbalance

import "sync"

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
	GetServers(mp map[string]interface{}) ([]*Service, error)
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
	Port       int
	Weight     int
	Status     InstanceStatus
	locker     sync.RWMutex
}
