package countdown

import (
	"context"
	"sync"
	"time"
)

// Countdown is a struct that encapsulates the countdown logic. It is directly usable as a context
type Countdown struct {
	sync.Mutex
	context.Context
	cancel   context.CancelFunc
	timeout  time.Duration
	timeleft time.Duration
}

// NewCountdown initialises a countdown context. t is the timeout to countdown to (based on calling Sub())
func NewCountdown(ctx context.Context, t time.Duration) *Countdown {
	c, cf := context.WithCancel(ctx)
	return &Countdown{
		Context:  c,
		cancel:   cf,
		timeout:  t,
		timeleft: t,
	}
}

// Sub subtracts the time t from timeout and if it's less than 0 cancels the context
func (c *Countdown) Sub(t time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.timeleft = c.timeleft - t

	if c.timeleft <= 0 {
		c.cancel()
	}
}
