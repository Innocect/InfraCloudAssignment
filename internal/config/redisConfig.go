package config

import "github.com/go-redis/redis"

type RedisConfig struct {
	Redis *redis.Client
}
