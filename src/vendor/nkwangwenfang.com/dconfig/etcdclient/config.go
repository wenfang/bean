package etcdclient

// Config etcd配置
type Config struct {
	Srvs []string `json:"srvs"`
	App  string   `json:"app"`
	Env  string   `json:"env"`
	Tag  string   `json:"tag"`
}
