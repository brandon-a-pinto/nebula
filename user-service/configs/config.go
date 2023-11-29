package configs

import "os"

type Config struct {
	DBDriver         string
	DBHost           string
	DBPort           string
	DBUsername       string
	DBPassword       string
	DBName           string
	GRPCServerPort   string
	BrokerServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:           os.Getenv("USER_DB_HOST"),
		DBPort:           os.Getenv("USER_DB_PORT"),
		DBUsername:       os.Getenv("USER_DB_USERNAME"),
		DBPassword:       os.Getenv("USER_DB_PASSWORD"),
		DBName:           os.Getenv("USER_DB_NAME"),
		GRPCServerPort:   os.Getenv("USER_GRPC_SERVER_PORT"),
		BrokerServerPort: os.Getenv("BROKER_WEB_SERVER_PORT"),
	}
}
