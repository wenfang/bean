package task

import (
	"golang.org/x/net/context"
)

// Factory 接口,创建Task
type Factory interface {
	// Factory名称
	Name() string
	// 创建Task
	CreateTask(int) Task
	// 获得所有的TaskID
	AllTaskIDs() []int
}

// Task 接口，Start启动任务
type Task interface {
	Start(context.Context) (context.CancelFunc, chan struct{}, error)
}