package seq

import (
	"testing"
)

var (
	expectID uint
	num      uint = 10000
	done          = make(chan struct{})
)

func do(seq *Seq, id uint, t *testing.T) {
	seq.Start(id)
	if id != expectID {
		t.Fatal("error seq")
	}
	expectID++
	seq.End(id)
	if id == num-1 {
		close(done)
	}
}

func TestSeq(t *testing.T) {
	var seq Seq
	var i uint

	for i = 0; i < num; i++ {
		id := seq.Next()
		go do(&seq, id, t)
	}
	<-done
}
