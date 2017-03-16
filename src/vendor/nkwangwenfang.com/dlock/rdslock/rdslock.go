package rdslock

import (
	"github.com/garyburd/redigo/redis"

	"nkwangwenfang.com/dlock"
	"nkwangwenfang.com/rds"
	"nkwangwenfang.com/util/uuid"
)

const (
	scriptLock string = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("PEXPIRE", KEYS[1], ARGV[2])
else
	local rv = redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
	if type(rv) == "table" then
		return 1
	else
		return 0
	end
end
`
	scriptUnlock string = `
if redis.call("get",KEYS[1]) == ARGV[1]
then
	return redis.call("del",KEYS[1])
else
	return 0
end
`
	// 锁缺省保有时间10s
	defalutHoldTime int = 10000
)

// RdsLock 基于Redis实现的分布式锁
type RdsLock struct {
	rds.Rdser
	key          string
	token        string
	holdTime     int
	lockScript   *redis.Script
	unlockScript *redis.Script
}

// Option 创建RdsLock的选项
type Option func(*RdsLock)

// HoldTime 设置RdsLock的持有时间
func HoldTime(holdTime int) Option {
	return func(l *RdsLock) {
		l.holdTime = holdTime
	}
}

// New 创建RdsLock锁
func New(r rds.Rdser, key string, options ...Option) dlock.Locker {
	l := &RdsLock{
		Rdser:        r,
		key:          key,
		holdTime:     defalutHoldTime,
		token:        uuid.New(),
		lockScript:   redis.NewScript(1, scriptLock),
		unlockScript: redis.NewScript(1, scriptUnlock),
	}

	for _, option := range options {
		option(l)
	}
	return l
}

// Lock 加锁，成功返回true,nil 加锁失败返回false,nil 脚本执行错误返回false,err
func (l *RdsLock) Lock() (bool, error) {
	reply, err := l.DoScript(l.lockScript, l.key, l.token, l.holdTime)
	if err != nil {
		return false, err
	}
	if rv, _ := redis.Int(reply, nil); rv == 0 {
		return false, nil
	}
	return true, nil
}

// Unlock 解锁，成功返回true,nil 解锁失败返回false,nil 脚本执行错误返回false,err
func (l *RdsLock) Unlock() (bool, error) {
	reply, err := l.DoScript(l.unlockScript, l.key, l.token)
	if err != nil {
		return false, err
	}
	if rv, _ := redis.Int(reply, nil); rv == 0 {
		return false, nil
	}
	return true, nil
}
