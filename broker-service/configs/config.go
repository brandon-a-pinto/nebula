package configs

import "os"

type Config struct {
	WebServerPort string
	UserGRPCPort  string
}

func LoadConfig() *Config {
	return &Config{
		WebServerPort: os.Getenv("BROKER_WEB_SERVER_PORT"),
		UserGRPCPort:  os.Getenv("USER_GRPC_SERVER_PORT"),
	}
}
