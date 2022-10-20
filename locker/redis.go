package locker

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisLocker 基于redis实现的Locker接口
type RedisLocker struct {
	rdb *redis.Conn
}

func (r *RedisLocker) Lock(ctx context.Context, key string, dur time.Duration) error {
	return nil
}

func (r *RedisLocker) Unlock(ctx context.Context, key string) error {
	return nil
}

func NewRedisLocker(rdb *redis.Conn) Locker {
	return &RedisLocker{rdb: rdb}
}