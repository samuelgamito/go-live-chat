package databases

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-live-chat/internal/configs"
)

type RedisClientInterface interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Process(ctx context.Context, cmd redis.Cmder) error
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}

type RedisClient struct {
	NotifyClientsRedis RedisClientInterface
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
