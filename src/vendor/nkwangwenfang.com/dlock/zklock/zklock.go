package zklock

import (
	"fmt"

	"github.com/samuel/go-zookeeper/zk"

	"nkwangwenfang.com/dlock"
)

type zkLock struct {
	conn *zk.Conn
	key  string
}

// New 创建基于zookeeper的Locker
func New(conn *zk.Conn, key string) dlock.Locker {
	return &zkLock{conn: conn, key: key}
}

func (zl *zkLock) Lock() (bool, error) {
	msg, err := zl.conn.Create(zl.key, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	fmt.Println(msg, err)
	return false, err
}

func (zl *zkLock) Unlock() (bool, error) {
	return true, nil
}
