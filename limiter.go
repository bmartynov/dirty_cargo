package cargo

import (
	"time"
	"sync/atomic"
)

type limiter struct {
	current int32
	maximum int32
	done    chan struct{}
}

func (l *limiter) loop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-l.done:
			return
		case <-ticker.C:
			atomic.StoreInt32(&l.current, 0)
		}
	}
}

func (l *limiter) Allow() bool {
	return atomic.AddInt32(&l.current, 1) < l.maximum
}

func (l *limiter) Stop() {
	close(l.done)
}

func newLimiter(perSec int32) Limiter {
	l := limiter{
		current: 0,
		maximum: perSec,
		done:    make(chan struct{}),
	}

	return &l
}
