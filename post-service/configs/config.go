package configs

import "os"

type Config struct {
	DBDriver       string
	DBHost         string
	DBPort         string
	DBUsername     string
	DBPassword     string
	DBName         string
	GRPCServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBDriver:       os.Getenv("POST_DB_DRIVER"),
		DBHost:         os.Getenv("POST_DB_HOST"),
		DBPort:         os.Getenv("POST_DB_PORT"),
		DBUsername:     os.Getenv("POST_DB_USERNAME"),
		DBPassword:     os.Getenv("POST_DB_PASSWORD"),
		DBName:         os.Getenv("POST_DB_NAME"),
		GRPCServerPort: os.Getenv("POST_GRPC_SERVER_PORT"),
	}
}
