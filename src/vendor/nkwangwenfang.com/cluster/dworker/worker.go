package dworker

import (
	"sort"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/node"
	"nkwangwenfang.com/cluster/oneshot"
	"nkwangwenfang.com/cluster/schedule"
	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/util/comp"
)

type worker struct {
	node.Node
	schedule.Scheduler
	task.Factory
	cancels     map[int]context.CancelFunc
	dones       map[int]chan struct{}
	lastTaskIDs []int
}

// NewWorker 创建Worker Oneshot，控制task执行
func NewWorker(node node.Node, schedule schedule.Scheduler, factory task.Factory) oneshot.Oneshot {
	return &worker{
		Node:      node,
		Scheduler: schedule,
		Factory:   factory,
		cancels:   make(map[int]context.CancelFunc),
		dones:     make(map[int]chan struct{}),
	}
}

// Do 作为Once的Do函数
func (w *worker) Do(c context.Context) error {
	// 获取调度信息
	info, err := w.GetInfo()
	if err != nil {
		log.Error("worker get info error", "worker", w.NodeID(), "error", err)
		return err
	}
	// 获取当前节点需执行的任务, 任务没有变化，直接返回
	taskIDs := info.GetTaskIDs(w.NodeID())
	if comp.IntsEqual(w.lastTaskIDs, taskIDs) {
		return nil
	}
	// 任务发生变化
	log.Info("task change", "new tasks", taskIDs, "worker", w.NodeID())
	// 启动新任务
	taskSet := make(map[int]struct{})
	for _, taskID := range taskIDs {
		taskSet[taskID] = struct{}{}
		if _, ok := w.cancels[taskID]; !ok {
			cancel, done, err := w.CreateTask(taskID).Start(c)
			if err != nil {
				log.Error("task start error", "taskID", taskID, "worker", w.NodeID(), "error", err)
				continue
			}
			w.cancels[taskID] = cancel
			w.dones[taskID] = done
			log.Info("task start", "taskID", taskID, "worker", w.NodeID())
		}
	}
	// 停止旧任务
	for _, taskID := range w.lastTaskIDs {
		if _, ok := taskSet[taskID]; !ok {
			w.cancels[taskID]()
			delete(w.cancels, taskID)
			<-w.dones[taskID]
			delete(w.dones, taskID)
			log.Info("task stop", "taskID", taskID, "worker", w.NodeID())
		}
	}
	// 设置lastTaskIDs
	w.lastTaskIDs = nil
	for taskID := range w.cancels {
		w.lastTaskIDs = append(w.lastTaskIDs, taskID)
	}
	sort.IntSlice(w.lastTaskIDs).Sort()
	return nil
}
