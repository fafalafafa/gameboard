package temporary

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// NewRedis creates a redis connection
func NewRedis() *redis.Client {
	fmt.Println("Redis connected")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}
