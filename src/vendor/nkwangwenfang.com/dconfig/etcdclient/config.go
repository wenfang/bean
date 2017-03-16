package etcdclient

// Config etcd配置
type Config struct {
	Servers []string `json:"servers"`
	App     string   `json:"app"`
}
