package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Db       int
}

func NewRedisClient(config RedisConfig) (*redis.Client, error) {
	if config.Host == "" {
		return nil, fmt.Errorf("redis config missing required field: Host")
	}
	if config.Port == "" {
		return nil, fmt.Errorf("redis config missing required field: Port")
	}
	if config.Username == "" {
		return nil, fmt.Errorf("redis config missing required field: Username")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("redis config missing required field: Password")
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
	}), nil
}

func GetKeyRedis(key string, value string) string {
	return fmt.Sprintf("%s:%s", key, value)
}
