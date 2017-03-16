package seq

import (
	"sync"
)

type sequencer struct {
	mu   sync.Mutex
	id   uint
	wait map[uint]chan uint
}

func (s *sequencer) start(id uint) {
	s.mu.Lock()
	if s.id == id {
		s.mu.Unlock()
		return
	}
	c := make(chan uint)
	if s.wait == nil {
		s.wait = make(map[uint]chan uint)
	}
	s.wait[id] = c
	s.mu.Unlock()
	<-c
}

func (s *sequencer) end(id uint) {
	s.mu.Lock()
	if s.id != id {
		panic("out of sync")
	}
	id++
	s.id = id
	if s.wait == nil {
		s.wait = make(map[uint]chan uint)
	}
	c, ok := s.wait[id]
	if ok {
		delete(s.wait, id)
	}
	s.mu.Unlock()
	if ok {
		c <- 1
	}
}

// Seq 顺序化控制
type Seq struct {
	sequencer

	mu sync.Mutex
	id uint
}

// Next 获得一个顺序编号
func (seq *Seq) Next() uint {
	seq.mu.Lock()
	id := seq.id
	seq.id++
	seq.mu.Unlock()
	return id
}

// Start 启动顺序任务
func (seq *Seq) Start(id uint) {
	seq.start(id)
}

// End 结束顺序任务
func (seq *Seq) End(id uint) {
	seq.end(id)
}
