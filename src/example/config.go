package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"nkwangwenfang.com/log"
)

// Config melon配置
type Config struct {
	MonitorAddr string `json:"monitor_addr"`
	LogLevel    string `json:"log_level"`
	LogFile     string `json:"log_file"`
	ErrFile     string `json:"err_file"`
	PidFile     string `json:"pid_file"`

	BindAddr string `json:"bind_addr"`
}

var configFile = flag.String("config", "conf/melon.conf", "config file name")

func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Error("load config file error", "filename", *configFile, "error", err)
		return nil, err
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		log.Error("config file unmarshal error", "error", err)
		return nil, err
	}
	return &cfg, nil
}
