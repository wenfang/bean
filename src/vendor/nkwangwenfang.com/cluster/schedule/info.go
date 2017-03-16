package schedule

import (
	"sort"
)

// Info 节点调度信息，NodeID到TaskID的映射表
type Info map[string][]int

// GetTaskIDs 获取对应nodeID的TaskID列表
func (info Info) GetTaskIDs(nodeID string) []int {
	taskIDs := info[nodeID]
	sort.IntSlice(taskIDs).Sort()
	return taskIDs
}
