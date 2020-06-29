package config

// Config type
type Config struct {
	Redis *RedisConfig
}

// RedisConfig creates a structure of redis configuration
type RedisConfig struct {
	Host string
	Port string
}

func redisConfig() *RedisConfig {
	return &RedisConfig{
		Host: "127.0.0.1",
		Port: "6379",
	}
}

// New creates an instance of config
func New() *Config {
	redis := redisConfig()
	return &Config{
		Redis: redis,
	}
}
