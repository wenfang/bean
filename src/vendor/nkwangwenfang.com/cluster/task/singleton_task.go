package task

import (
	"errors"

	"golang.org/x/net/context"
)

var (
	// ErrIsRunning task正在执行
	ErrIsRunning = errors.New("task is running")
)

type singletonTask struct {
	task      Task
	isRunning bool
}

// WithSingleton 封装task，使其做到单例执行
func WithSingleton(task Task) Task {
	return &singletonTask{task: task}
}

func (st *singletonTask) Start(ctx context.Context) (context.CancelFunc, chan struct{}, error) {
	if st.isRunning {
		return nil, nil, ErrIsRunning
	}

	cancel, done, err := st.task.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	st.isRunning = true
	return cancel, done, nil
}
