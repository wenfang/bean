package oneshot

import (
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/log"
)

type timerOneshot struct {
	oneshot Oneshot
	d       time.Duration
}

// WithTimer 延迟oneshot到d时间后执行
func WithTimer(oneshot Oneshot, d time.Duration) Oneshot {
	return &timerOneshot{oneshot: oneshot, d: d}
}

func (t *timerOneshot) Do(ctx context.Context) error {
	time.AfterFunc(t.d, func() {
		if err := t.oneshot.Do(ctx); err != nil {
			log.Error("timer once do failed", "error", err)
		}
	})

	return nil
}
