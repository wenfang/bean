package batcher

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBatches(t *testing.T) {
	var iters uint32

	rand.Seed(time.Now().Unix())
	bulk := rand.Intn(10)
	b := New(10*time.Millisecond, uint(bulk), func(params []interface{}) {
		atomic.AddUint32(&iters, uint32(len(params)))
	})

	len := rand.Intn(1000)
	wg := &sync.WaitGroup{}
	for i := 0; i < len; i++ {
		wg.Add(1)
		go func() {
			b.Run(nil)
			wg.Done()
		}()
	}
	wg.Wait()

	if iters != uint32(len) {
		t.Error("Wrong number of iters:", iters)
	}
}

func TestZero(t *testing.T) {
	b := New(0, 1, func(params []interface{}) {})
	for i := 0; i < 100; i++ {
		b.Run(nil)
	}
}
