package locker

import (
	"context"
	"time"
)

// Locker locker通用接口
type Locker interface {
	Lock(ctx context.Context, key string, dur time.Duration) error
	Unlock(ctx context.Context, key string)
	// SetLoop 设置自旋次数(非必要，不设置则默认值3次)
	SetLoop(loop int)
	// SetLoopWaitTime 设置每次等待时间，单位:ms(非必要，不设置则默认值100ms)
	SetLoopWaitTime(dur int)
}

type Option func(locker Locker)

const (
	DefaultLockerLoop = 3
	DefaultLockerLoopWaitTime = 100
)
