package dworker

import (
	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/node"
	"nkwangwenfang.com/cluster/oneshot"
	"nkwangwenfang.com/cluster/schedule"
	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/util/comp"
)

// leader 集群中处于leader状态
type leader struct {
	node.Node
	schedule.Scheduler
	task.Factory
	// 上次获得的所有NodeID
	lastNodeIDs []string
}

// NewLeader 创建leader Once
func NewLeader(node node.Node, schedule schedule.Scheduler, factory task.Factory) oneshot.Oneshot {
	return &leader{
		Node:      node,
		Scheduler: schedule,
		Factory:   factory,
	}
}

func rebalance(nodeIDs []string, taskIDs []int) schedule.Info {
	info := make(schedule.Info)
	nodeNum := len(nodeIDs)
	if nodeNum != 0 {
		// rebalance all tasks
		for idx, taskID := range taskIDs {
			nodeID := nodeIDs[idx%nodeNum]
			info[nodeID] = append(info[nodeID], taskID)
		}
	}
	return info
}

// Do 作为Once的Do函数
func (l *leader) Do(c context.Context) error {
	// 获取所有节点ID
	nodeIDs, err := l.AllNodeIDs()
	if err != nil {
		log.Error("leader get all node id error", "leader", l.NodeID(), "error", err)
		return err
	}
	// 节点没变化，直接返回
	if comp.StringsEqual(l.lastNodeIDs, nodeIDs) {
		return nil
	}
	// 重新平衡节点及任务
	info := rebalance(nodeIDs, l.AllTaskIDs())
	if err = l.SetInfo(info); err != nil {
		log.Error("leader set info error", "leader", l.NodeID(), "error", err)
		return err
	}
	log.Info("schedule", "leader", l.NodeID(), "info change", info)
	// 存储节点ID
	l.lastNodeIDs = nodeIDs
	return nil
}
