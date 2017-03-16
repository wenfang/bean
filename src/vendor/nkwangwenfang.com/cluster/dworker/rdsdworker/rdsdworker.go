package rdsdworker

import (
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/dworker"
	"nkwangwenfang.com/cluster/master/rdsmaster"
	"nkwangwenfang.com/cluster/node/rdsnode"
	"nkwangwenfang.com/cluster/oneshot"
	"nkwangwenfang.com/cluster/schedule/rdsschedule"
	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/rds"
)

const defaultLoopTime time.Duration = 3e9

// Bundle 集合dworker服务
type Bundle struct {
	node *rdsnode.RdsNode

	nodeTask   task.Task
	leaderTask task.Task
	workerTask task.Task

	cancel context.CancelFunc

	nodeDone   chan struct{}
	leaderDone chan struct{}
	workerDone chan struct{}
}

// New 创建基于redis的dworker Once
func New(r rds.Rdser, name string, factory task.Factory) *Bundle {
	node := rdsnode.New(r, name)
	scheduler := rdsschedule.New(r, name)
	master := rdsmaster.New(r, name)

	nodeTask := task.WithLoop(node, defaultLoopTime)
	leaderTask := task.WithLoop(oneshot.WithMaster(dworker.NewLeader(node, scheduler, factory), master), defaultLoopTime)
	workerTask := task.WithLoop(dworker.NewWorker(node, scheduler, factory), defaultLoopTime)

	return &Bundle{
		node:       node,
		nodeTask:   nodeTask,
		leaderTask: leaderTask,
		workerTask: workerTask,
	}
}

// Start 启动bundle
func (b *Bundle) Start(c context.Context) error {
	var (
		nodeCancel, leaderCancel context.CancelFunc
		err                      error
	)

	ctx, cancel := context.WithCancel(c)

	nodeCancel, b.nodeDone, err = b.nodeTask.Start(ctx)
	if err != nil {
		log.Error("node task start error", "error", err)
		return err
	}

	leaderCancel, b.leaderDone, err = b.leaderTask.Start(ctx)
	if err != nil {
		log.Error("leader task start error", "error", err)
		nodeCancel()
		<-b.nodeDone
		return err
	}

	_, b.workerDone, err = b.workerTask.Start(ctx)
	if err != nil {
		log.Error("worker task start error", "error", err)
		nodeCancel()
		leaderCancel()
		<-b.nodeDone
		<-b.leaderDone
		return err
	}
	b.cancel = cancel
	return nil
}

// Stop 停止bundle
func (b *Bundle) Stop() {
	b.cancel()
	<-b.nodeDone
	b.node.Remove()
	<-b.leaderDone
	<-b.workerDone
}
