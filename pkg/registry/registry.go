package registry

type Registry interface {
	Register(name, nodeID, addr string, tags ...string) error
	Unregister(name, nodeID string) error
}
