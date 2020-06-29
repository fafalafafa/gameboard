package datasource

import (
	"github.com/fafalafafa/gameboard/datasource/temporary"
	"github.com/go-redis/redis/v8"
)

// DataSource is the dataType
type DataSource struct {
	Redis *redis.Client
}

// NewTemporary creates temporary storage connectivity like redis and other cache solutions
func newRedis() *redis.Client {
	temporary := temporary.NewRedis()
	return temporary
}

// InitDS defines the whole datasource
func InitDS() *DataSource {
	redisConn := newRedis()
	return &DataSource{
		Redis: redisConn,
	}
}
