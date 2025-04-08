package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// ServerHost string `envconfig:"HEX_SERVER_HOST" default:"127.0.0.1"`
	// ServerPort string `envconfig:"HEX_SERVER_PORT" default:"5000"`

	// DBHost     string `envconfig:"HEX_DB_HOST" default:"localhost"`
	// DBPort     string `envconfig:"HEX_DB_PORT" default:"5432"`
	// DBUsername string `envconfig:"HEX_DB_USERNAME" default:"postgres"`
	// DBPassword string `envconfig:"HEX_DB_PASSWORD" default:"postgres"`
	// DBName     string `envconfig:"HEX_DB_NAME" default:"yourdb"`
	// DBSSLMode  string `envconfig:"HEX_DB_SSLMODE" default:"disable"`

	// RedisHost     string `envconfig:"HEX_REDIS_HOST" default:"localhost"`
	// RedisPort     string `envconfig:"HEX_REDIS_PORT" default:"6379"`
	// RedisPassword string `envconfig:"HEX_REDIS_PASSWORD" default:""`
	// RedisDb       int    `envconfig:"HEX_REDIS_DB" default:"0"`

	Server   Server
	Database Database
	Redis    Redis
}

type Server struct {
	Host string `envconfig:"HEX_SERVER_HOST" default:"127.0.0.1"`
	Port string `envconfig:"HEX_SERVER_PORT" default:"5000"`
}

type Database struct {
	Host     string `envconfig:"HEX_DB_HOST" default:"127.0.0.1"`
	Port     string `envconfig:"HEX_DB_PORT" default:"5432"`
	Username string `envconfig:"HEX_DB_USERNAME" default:"postgres"`
	Password string `envconfig:"HEX_DB_PASSWORD" default:"1q2w3e4r5t"`
	Name     string `envconfig:"HEX_DB_NAME" default:"hexdb_server"`
	SSLMode  string `envconfig:"HEX_DB_SSLMODE" default:"disable"`
}

type Redis struct {
	Host     string `envconfig:"HEX_REDIS_HOST" default:"localhost"`
	Port     string `envconfig:"HEX_REDIS_PORT" default:"6379"`
	Password string `envconfig:"HEX_REDIS_PASSWORD" default:""`
	Db       int    `envconfig:"HEX_REDIS_DB" default:"2"`
}

func Load() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load env variable into config struct: %w", err)
	}

	return &cfg, nil
}

// func LoadConfig() *Config {
// 	cfg := &Config{}
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}
// 	err = env.Parse(cfg)
// 	if err != nil {
// 		panic("Error parsing environment variables")
// 	}

// 	return cfg
// }
