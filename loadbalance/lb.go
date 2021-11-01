package loadbalance

import "container/list"

type Lb interface {
	Description() string
	GetServers() []*Service
	GetInstance(serviceName string) *Instance
}

type RobinLb struct {
	instances map[string]*list.List
	d         Discovery
}

func NewRobinLb(dis Discovery) *RobinLb {
	return &RobinLb{
		d:         dis,
		instances: make(map[string]*list.List),
	}
}

func (r *RobinLb) Description() string {
	return "roubin load balance"
}

func (r *RobinLb) GetServers() ([]*Service, error) {
	return r.d.GetServers()
}

func (r *RobinLb) GetInstance(serviceName string) *Instance {
	insList := r.instances[serviceName]
	if insList != nil {
		return r.getInstance(serviceName)
	}
	servers, err := r.GetServers()
	if err != nil {
		return nil
	}
	var ins []*Instance
	for _, s := range servers {
		if s.Name == serviceName {
			ins = s.instances
			break
		}
	}
	if len(ins) > 0 {
		l := list.New()
		for _, s := range ins {
			l.PushBack(s)
		}
		r.instances[serviceName] = l
	}
	return r.getInstance(serviceName)
}

func (r *RobinLb) getInstance(serviceName string) *Instance {
	instances := r.instances[serviceName]
	if nil != instances {
		next := instances.Front()
		value := next.Value
		instances.Remove(next)
		instances.PushBack(value)
		return value.(*Instance)
	}
	return nil
}
