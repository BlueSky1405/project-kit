package locker

import (
	"context"
	"time"
)

// Locker locker通用接口
type Locker interface {
	Lock(ctx context.Context, key string, dur time.Duration) error
	Unlock(ctx context.Context, key string) error
}
