package log

// Config log的配置结构
type Config struct {
	LogLevel string `json:"log_level"`
	LogFile  string `json:"log_file"`
	ErrFile  string `json:"err_file"`
}
