package config

import (
	"encoding/json"
	"fmt"
)

// Config struct for Warden service
type WardenConfig struct {
	Config *Config `json:"config"`
}

type Config struct {
	Env      string          `json:"env"`
	GRPCPort int             `json:"grpc_port"`
	SSODB    *DatabaseConfig `json:"sso_db"`
	Broker   *BrokerConfig   `json:"broker"`
}

// Database connection configuration
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"sslmode"`
}

// Message broker connection configuration
type BrokerConfig struct {
	Type       string `json:"type"`
	Address    string `json:"address"`
	Datacenter string `json:"datacenter"`
}

// MapConfig maps bytes of config into Config struct
func MapConfig(data []byte) (*Config, error) {
	c := new(WardenConfig)
	if err := json.Unmarshal(data, c); err != nil {
		return nil, fmt.Errorf("error map config to struct. %w", err)
	}

	return c.Config, nil
}

//	{
//		"env":"local",
//		"grpc_port": 3003,
//			"database":{
//			"address":"localhost:5430"
//		},
//		"broker":{
//			"type":"nats",
//			"address":"localhost:4555"
//		}
//	}
var DefaultConfig *Config = &Config{
	Env:      "local",
	GRPCPort: 3003,
	SSODB: &DatabaseConfig{
		Host:     "localhost",
		Port:     "5430",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	},
	Broker: &BrokerConfig{
		Type:       "nats",
		Address:    "localhost:4555",
		Datacenter: "dc",
	},
}
