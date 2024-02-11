package events

type BrokerConfig interface {
	Datacenter() string
	Host() string
}
