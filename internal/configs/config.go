package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Port               string
	OpenChatMongoDB    *MongoDBConfig
	NotifyClientsRedis *RedisConfig
}

func NewEnvConfig() *Config {
	return &Config{
		Port: getEnv("port", "8080"),
		OpenChatMongoDB: &MongoDBConfig{
			User:     getEnv("mongoUser", "admin"),
			Pass:     getEnv("mongoPass", "admin"),
			Host:     getEnv("mongoHost", "localhost"),
			Protocol: getEnv("mongoProtocol", "mongodb"),
			Port:     getEnv("mongoPort", "27017"),
			Database: getEnv("mongoDatabase", "Develop"),
		},
		NotifyClientsRedis: &RedisConfig{
			Addr:     getEnv("redisAddr", "localhost:6379"),
			Password: getEnv("redisPassword", ""),
			DB:       getEnvAsInt("redisDB", "0"),
			Protocol: getEnvAsInt("redisProtocol", "2"),
		},
	}
}

func getEnvAsInt(key string, defaultVal string) int {
	val := getEnv(key, defaultVal)
	i, err := strconv.Atoi(val)

	if err != nil {
		panic(err)
	}

	return i
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
