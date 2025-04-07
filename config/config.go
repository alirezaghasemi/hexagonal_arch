package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string `envconfig:"HEX_SERVER_HOST" envDefault:"127.0.0.1"`
	ServerPort string `envconfig:"HEX_SERVER_PORT" envDefault:"5000"`

	DBHost     string `envconfig:"HEX_DB_HOST" envDefault:"localhost"`
	DBPort     string `envconfig:"HEX_DB_PORT" envDefault:"5432"`
	DBUsername string `envconfig:"HEX_DB_USERNAME" envDefault:"postgres"`
	DBPassword string `envconfig:"HEX_DB_PASSWORD" envDefault:"postgres"`
	DBName     string `envconfig:"HEX_DB_NAME" envDefault:"yourdb"`
	DBSSLMode  string `envconfig:"HEX_DB_SSLMODE" envDefault:"disable"`

	RedisHost     string `envconfig:"HEX_REDIS_HOST" envDefault:"localhost"`
	RedisPort     string `envconfig:"HEX_REDIS_PORT" envDefault:"6379"`
	RedisPassword string `envconfig:"HEX_REDIS_PASSWORD" envDefault:""`
	RedisDb       int    `envconfig:"HEX_REDIS_DB" envDefault:"0"`

	// Server   Server
	// Database Database
	// Redis    Redis
}

// type Server struct {
// 	Host string `envconfig:"HEX_SERVER_HOST" envDefault:"127.0.0.1"`
// 	Port string `envconfig:"HEX_SERVER_PORT" envDefault:"5000"`
// }

// type Database struct {
// 	Host     string `envconfig:"HEX_DB_HOST" envDefault:"localhost"`
// 	Port     string `envconfig:"HEX_DB_PORT" envDefault:"5432"`
// 	Username string `envconfig:"HEX_DB_USERNAME" envDefault:"postgres"`
// 	Password string `envconfig:"HEX_DB_PASSWORD" envDefault:"postgres"`
// 	Name     string `envconfig:"HEX_DB_NAME" envDefault:"yourdb"`
// 	SSLMode  string `envconfig:"HEX_DB_SSLMODE" envDefault:"disable"`
// }

// type Redis struct {
// 	Host     string `envconfig:"HEX_REDIS_HOST" envDefault:"localhost"`
// 	Port     string `envconfig:"HEX_REDIS_PORT" envDefault:"6379"`
// 	Password string `envconfig:"HEX_REDIS_PASSWORD" envDefault:""`
// 	Db       int    `envconfig:"HEX_REDIS_DB" envDefault:"0"`
// }

func LoadConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	err = env.Parse(cfg)
	if err != nil {
		panic("Error parsing environment variables")
	}

	return cfg
}
