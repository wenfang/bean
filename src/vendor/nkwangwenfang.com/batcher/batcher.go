// Package batcher 实现批量处理
package batcher

import (
	"sync"
	"time"
)

// work代表一个任务, done控制执行是否完成
type work struct {
	param interface{}
	done  chan struct{}
}

// Batcher 批量执行的结构
type Batcher struct {
	sync.Mutex
	works []*work

	timeout time.Duration
	cnt     uint
	do      func([]interface{})
}

// New 创建一个批量执行的结构,timeout为最长等待时间,cnt为汇聚多少任务后执行,do为要执行的函数
func New(timeout time.Duration, cnt uint, do func([]interface{})) *Batcher {
	b := &Batcher{
		works:   make([]*work, 0, cnt),
		timeout: timeout,
		cnt:     cnt,
		do:      do,
	}
	go b.timer()
	return b
}

func (b *Batcher) timer() {
	for range time.Tick(b.timeout) {
		b.Lock()
		if len(b.works) != 0 {
			go b.batch(b.works)
			b.works = make([]*work, 0, b.cnt)
		}
		b.Unlock()
	}
}

// batch 批量处理提交的work
func (b *Batcher) batch(works []*work) {
	params := make([]interface{}, 0, b.cnt)
	for _, w := range works {
		params = append(params, w.param)
	}

	b.do(params)

	for _, w := range works {
		close(w.done)
	}
}

// Run 执行一个任务,会停在这里，直到任务执行完成
func (b *Batcher) Run(param interface{}) {
	if b.timeout == 0 || b.cnt <= 1 {
		b.do([]interface{}{param})
		return
	}

	w := &work{
		param: param,
		done:  make(chan struct{}),
	}
	b.submit(w)

	<-w.done
}

func (b *Batcher) submit(w *work) {
	b.Lock()
	defer b.Unlock()

	b.works = append(b.works, w)
	if len(b.works) == int(b.cnt) {
		go b.batch(b.works)
		b.works = make([]*work, 0, b.cnt)
	}
}
