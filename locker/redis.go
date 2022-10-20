package locker

import (
	"context"
	"github.com/BlueSky1405/project-kit/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

const RedisKeyLockerPrefix = "locker_"

// RedisLocker 基于redis实现的Locker接口
type RedisLocker struct {
	// 自旋次数
	loop   int
	loopWt int
	rdb    *redis.Client
	logger log.Logger
}

func (r *RedisLocker) SetLoop(loop int) {
	r.loop = loop
}

func (r *RedisLocker) SetLoopWaitTime(dur int) {
	r.loopWt = dur
}

func (r *RedisLocker) Lock(ctx context.Context, key string, dur time.Duration) error {
	for i := 0; i < r.loop; i++ {
		result, err := r.rdb.SetNX(ctx, key, 1, dur).Result()
		if err != nil {
			return errors.Wrapf(err, "RedisLocker Lock fail, key:%s", key)
		}
		if result {
			return nil
		}
		// 直接睡眠，会导致协程上下文切换，待优化
		time.Sleep(time.Millisecond * time.Duration(r.loopWt))
	}
	return errors.New("RedisLocker obtain lock fail")
}

func (r *RedisLocker) Unlock(ctx context.Context, key string) {
	// TODO 错误重试机制
	if err := r.rdb.Del(ctx, key).Err(); err != nil {
		r.logger.ErrorW("RedisLocker release lock fail", "key", key)
	}
}

func NewRedisLocker(rdb *redis.Client, logger log.Logger, options ...Option) Locker {
	l := &RedisLocker{rdb: rdb, logger: logger}

	for _, opt := range options {
		opt(l)
	}

	if l.loop == 0 {
		l.loop = DefaultLockerLoop
	}
	if l.loopWt == 0 {
		l.loopWt = DefaultLockerLoopWaitTime
	}

	return l
}

func WithLoop(loop int) Option {
	return func(locker Locker) {
		locker.SetLoop(loop)
	}
}

func WithLoopWaitTime(waitTime int) Option {
	return func(locker Locker) {
		locker.SetLoopWaitTime(waitTime)
	}
}
