package consul

import (
	"fmt"
	"net"
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

type Consul struct {
	agent *consul.Client
}

func (c *Consul) Register(name, nodeID, addr string, tags ...string) error {
	host, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("error parse host string. %w", err)
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return fmt.Errorf("error parse port value. %w", err)
	}

	return c.agent.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID:      name + "-" + nodeID,
		Name:    name,
		Port:    port,
		Address: host,

		Check: &consul.AgentServiceCheck{
			TCP:                            addr,
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "20s",
			Status:                         "passing",
		},
	})
}

func (c *Consul) Client() *consul.Client {
	return c.agent
}

func (c *Consul) Unregister(name, nodeID string) error {
	return c.agent.Agent().ServiceDeregister(name + "-" + nodeID)
}

func NewConsul(addr string, dc string, token string, tags ...string) (*Consul, error) {
	consulAgent, err := consul.NewClient(&consul.Config{
		Address:    addr,
		Datacenter: dc,
		Token:      token,
	})
	if err != nil {
		return nil, fmt.Errorf("error connect to consul. %w", err)
	}

	return &Consul{
		agent: consulAgent,
	}, nil
}
