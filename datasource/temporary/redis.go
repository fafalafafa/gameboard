package temporary

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// NewRedis creates a redis connection
func NewRedis() *redis.Client {
	fmt.Println("Redis connected")
	conn := os.Getenv("REDIS_CONNECTION")

	if conn == "" {
		conn = "localhost:6379"
	}

	fmt.Println(conn)
	rdb := redis.NewClient(&redis.Options{
		Addr: conn,
	})

	return rdb
}
