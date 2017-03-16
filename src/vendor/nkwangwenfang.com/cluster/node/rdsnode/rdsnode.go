package rdsnode

import (
	"fmt"
	"sort"

	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"

	"nkwangwenfang.com/log"
	"nkwangwenfang.com/rds"
	"nkwangwenfang.com/util/uuid"
)

const (
	clusterKey string = "%s_cluster"
	nodeKey    string = "%s_node_%s"

	// nodeKey的缺省保有时间10s
	defaultHoldTime int = 10000
)

// RdsNode 结构
type RdsNode struct {
	rds.Rdser
	clusterID string
	nodeID    string
	holdTime  int
}

// Option 设置集群选项
type Option func(*RdsNode)

// HoldTime 设置节点Key在redis中的保有时间
func HoldTime(holdTime int) Option {
	return func(n *RdsNode) {
		n.holdTime = holdTime
	}
}

// New 创建redis集群节点
func New(r rds.Rdser, name string, options ...Option) *RdsNode {
	n := &RdsNode{
		Rdser:     r,
		clusterID: fmt.Sprintf(clusterKey, name),
		nodeID:    fmt.Sprintf(nodeKey, name, uuid.New()),
		holdTime:  defaultHoldTime,
	}

	for _, option := range options {
		option(n)
	}
	return n
}

// Do 作为Once的Do函数
func (n *RdsNode) Do(_ context.Context) error {
	if _, err := n.DoCmd("SET", n.nodeID, n.nodeID, "PX", n.holdTime); err != nil {
		log.Error("update nodeID error", "nodeID", n.nodeID, "error", err)
	}

	if _, err := n.DoCmd("SADD", n.clusterID, n.nodeID); err != nil {
		log.Error("add nodeID error", "clusterID", n.clusterID, "nodeID", n.nodeID, "error", err)
	}
	return nil
}

// Remove 从集群中移除节点
func (n *RdsNode) Remove() {
	n.remove(n.nodeID)
}

// ClusterID 返回集群ID
func (n *RdsNode) ClusterID() string {
	return n.clusterID
}

// NodeID 返回节点ID
func (n *RdsNode) NodeID() string {
	return n.nodeID
}

func (n *RdsNode) check(nodeID string) bool {
	reply, err := n.DoCmd("GET", nodeID)
	if err != nil {
		log.Error("get nodeID error", "nodeID", nodeID, "error", err)
		return true
	}
	if _, err := redis.String(reply, err); err != nil {
		return false
	}
	return true
}

func (n *RdsNode) remove(nodeID string) {
	if _, err := n.DoCmd("SREM", n.clusterID, nodeID); err != nil {
		log.Error("remove nodeID from clusterID error", "clusterID", n.clusterID, "nodeID", nodeID, "error", err)
	}
	if _, err := n.DoCmd("DEL", nodeID); err != nil {
		log.Error("remove nodeID error", "nodeID", nodeID, "error", err)
	}
}

// AllNodeIDs 返回所有节点ID
func (n *RdsNode) AllNodeIDs() ([]string, error) {
	nodeIDs, err := redis.Strings(n.DoCmd("SMEMBERS", n.clusterID))
	if err != nil {
		log.Error("get nodeIDs error", "clusterID", n.clusterID, "error", err)
		return nil, err
	}

	var validNodeIDs []string
	for _, nodeID := range nodeIDs {
		if ok := n.check(nodeID); !ok {
			log.Warning("nodeID not found, remove from group", "nodeID", nodeID)
			n.remove(nodeID)
			continue
		}
		validNodeIDs = append(validNodeIDs, nodeID)
	}
	nodeIDs = validNodeIDs

	sort.StringSlice(nodeIDs).Sort()
	return nodeIDs, nil
}
