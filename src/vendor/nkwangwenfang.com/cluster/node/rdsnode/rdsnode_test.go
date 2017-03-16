package rdsnode

import (
	"math/rand"
	"testing"
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/rds/codis"
)

var codisConfig = codis.Config{
	Addrs:          []string{"127.0.0.1:6379"},
	MaxIdle:        20,
	MaxActive:      50,
	ConnectTimeout: 3000,
	ReadTimeout:    3000,
	WriteTimeout:   3000,
	IdleTimeout:    30000,
	Wait:           true,
}

func TestRdsNode(t *testing.T) {
	r := codis.New(codisConfig)
	defer r.Close()

	num := rand.Intn(101)
	nodes := make([]*RdsNode, num)
	cancels := make([]context.CancelFunc, num)
	dones := make([]chan struct{}, num)
	nodeMap := make(map[string]struct{})
	for i := 0; i < num; i++ {
		nodes[i] = New(r, "test")
		nodeMap[nodes[i].NodeID()] = struct{}{}
		cancels[i], dones[i], _ = task.WithLoop(nodes[i], 3e9).Start(context.Background())
	}

	time.Sleep(5e9)
	nodes[0].AllNodeIDs()

	for i := 0; i < num; i++ {
		delete(nodeMap, nodes[i].NodeID())
		cancels[i]()
		<-dones[i]
		nodes[i].Remove()
	}
}
