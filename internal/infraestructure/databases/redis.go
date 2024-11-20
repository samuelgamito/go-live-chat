package databases

import (
	"github.com/redis/go-redis/v9"
	"go-live-chat/internal/configs"
)

type RedisClient struct {
	NotifyClientsRedis *redis.Client
}

func NewRedisClient(config *configs.Config) *RedisClient {
	return &RedisClient{
		NotifyClientsRedis: redis.NewClient(&redis.Options{
			Username: config.NotifyClientsRedis.Username,
			Addr:     config.NotifyClientsRedis.Addr,
			Password: config.NotifyClientsRedis.Password,
			DB:       config.NotifyClientsRedis.DB,
			Protocol: config.NotifyClientsRedis.Protocol,
		}),
	}
}
