package rdsmaster

import (
	"fmt"

	"nkwangwenfang.com/cluster/master"
	"nkwangwenfang.com/dlock"
	"nkwangwenfang.com/dlock/rdslock"
	"nkwangwenfang.com/rds"
)

const (
	// 锁的缺省持有时间10s
	defaultHoldTime int = 10000

	masterKey string = "%s_master"
)

// RdsMaster 结构
type RdsMaster struct {
	dlock.Locker
	name     string
	holdTime int
}

// Option Master结构的选项
type Option func(*RdsMaster)

// HoldTime 设置master角色保有时间
func HoldTime(holdTime int) Option {
	return func(m *RdsMaster) {
		m.holdTime = holdTime
	}
}

// New 创建Master结构
func New(r rds.Rdser, name string, options ...Option) master.Master {
	m := &RdsMaster{
		name:     name,
		holdTime: defaultHoldTime,
	}
	for _, option := range options {
		option(m)
	}

	key := fmt.Sprintf(masterKey, name)
	m.Locker = rdslock.New(r, key, rdslock.HoldTime(m.holdTime))
	return m
}

// CheckMaster 检查是否为Master
func (m *RdsMaster) CheckMaster() (bool, error) {
	return m.Lock()
}
