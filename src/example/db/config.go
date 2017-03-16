package db

// Config db的配置结构
type Config struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	DBName string `json:"dbname"`
}
