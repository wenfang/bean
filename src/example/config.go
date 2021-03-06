package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"nkwangwenfang.com/kit/monitor"
	"nkwangwenfang.com/log"
)

// Config melon配置
type Config struct {
	Monitor monitor.Config `json:"monitor"`
	Log     log.Config     `json:"log"`
	PidFile string         `json:"pid_file"`

	BindAddr string `json:"bind_addr"`
}

var configFile = flag.String("config", "conf/example.conf", "config file name")

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
