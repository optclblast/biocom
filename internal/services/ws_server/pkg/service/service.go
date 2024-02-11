package service

type Service struct {
	Name     string
	Metadata map[string]any
	Version  string
	Address  string
	Protocol string
}
