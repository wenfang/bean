package task

import (
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/master"
	"nkwangwenfang.com/log"
)

type masterTask struct {
	task     Task
	interval time.Duration
	master   master.Master
	isMaster bool

	cancel context.CancelFunc
	done   chan struct{}
}

// WithMaster 包装Task使其具有master功能
func WithMaster(task Task, interval time.Duration, master master.Master) Task {
	return &masterTask{task: task, interval: interval, master: master}
}

func (mt *masterTask) Start(c context.Context) (context.CancelFunc, chan struct{}, error) {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(c)

	go func() {
		defer close(done)

		// 先执行一次
		mt.do(ctx)

		tick := time.Tick(mt.interval)
		for {
			select {
			case <-tick:
				mt.do(ctx)
			case <-ctx.Done():
				if mt.isMaster {
					mt.cancel()
					<-mt.done
				}
				return
			}
		}
	}()
	return cancel, done, nil
}

func (mt *masterTask) do(ctx context.Context) {
	// 检查是否仍是master
	isMaster, err := mt.master.CheckMaster()
	if err != nil {
		log.Error("check master error", "error", err)
		return
	}
	// 没有变化，直接返回
	if mt.isMaster == isMaster {
		return
	}
	if isMaster {
		// 如果变为master了，启动Task执行，如果启动失败，下次继续争抢
		var err error
		mt.cancel, mt.done, err = mt.task.Start(ctx)
		if err != nil {
			log.Error("task start error", "error", err)
			return
		}
	} else {
		// 如果不是Master了，但以前是，停止task
		mt.cancel()
		<-mt.done
	}
	mt.isMaster = isMaster
}
