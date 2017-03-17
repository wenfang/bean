package uuid

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"nkwangwenfang.com/utils/ipaddr"
)

const (
	// token 起始时间为2015-1-1 00:00:00 UTC
	startTime = 1420070400
)

var seq uint32

// New 创建新的唯一uuid
func New() string {
	localIP, err := ipaddr.LocalIPv4()
	if err != nil {
		localIP = "127.0.0.1"
	}
	return fmt.Sprintf("%s_%d_%d_%d",
		localIP,
		os.Getpid(),
		time.Now().Unix()-startTime,
		atomic.AddUint32(&seq, 1),
	)
}
