package node

// Node 接口
type Node interface {
	ClusterID() string
	NodeID() string
	AllNodeIDs() ([]string, error)
}
