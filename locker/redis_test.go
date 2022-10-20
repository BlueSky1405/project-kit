package locker

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisLocker(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	})
	locker := NewRedisLocker(rdb, WithLoop(DefaultLockerLoop), WithLoopWaitTime(DefaultLockerLoopWaitTime))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	key := "深圳的锁"
	defer func() {
		_ = locker.Unlock(ctx, key)
	}()

	err := locker.Lock(ctx, key, 5*time.Second)
	require.Nil(t, err)

	err = locker.Lock(ctx, key, 3*time.Second)
	require.Equal(t, "RedisLocker obtain lock fail", err.Error())
}
