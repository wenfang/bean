package expvar

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	cnt uint64
}

func (c *Counter) Add(delta uint64) {
	atomic.AddUint64(&c.cnt, delta)
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d", c.cnt)
}

func NewCounter(name string) *Counter {
	c := &Counter{cnt: 0}
	publish(name, c)
	return c
}
