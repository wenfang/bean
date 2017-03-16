package schedule

// Scheduler 获取或设置节点调度信息
type Scheduler interface {
	// 获得调度信息
	GetInfo() (Info, error)
	// 设置调度信息
	SetInfo(info Info) error
}
