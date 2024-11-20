package configs

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
	Protocol int
}
