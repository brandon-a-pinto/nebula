package configs

import "os"

type Config struct {
	DBHost         string
	DBName         string
	GRPCServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:         os.Getenv("LOGGER_DB_HOST"),
		DBName:         os.Getenv("LOGGER_DB_NAME"),
		GRPCServerPort: os.Getenv("LOGGER_GRPC_SERVER_PORT"),
	}
}
