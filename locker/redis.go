package locker

import "github.com/go-redis/redis/v8"

// RedisLocker 基于redis实现的Locker接口
type RedisLocker struct {
	rdb *redis.Conn
}

func NewRedisLocker(rdb *redis.Conn) *RedisLocker {
	return &RedisLocker{rdb: rdb}
}