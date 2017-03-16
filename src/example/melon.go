package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "nkwangwenfang.com/kit/expvar"
	"nkwangwenfang.com/kit/monitor"
	// 注释压制golint
	_ "nkwangwenfang.com/kit/pprof"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/util/pidfile"

	"xiaojukeji.com/melon/db"
)

var GlobalDB *db.StatusDB

func main() {
	flag.Parse()
	// 解析配置文件
	cfg, err := loadConfig()
	if err != nil {
		return
	}
	// 设置日志
	log.Init(cfg.LogFile, cfg.ErrFile, cfg.LogLevel)
	// 创建pid文件
	if cfg.PidFile != "" {
		if err := pidfile.Save(cfg.PidFile); err != nil {
			log.Error("save pid file error", "pid_file", cfg.PidFile)
		}
		defer pidfile.Remove(cfg.PidFile)
	}
	// 启动monitor
	if cfg.MonitorAddr != "" {
		go func() {
			if err := monitor.ListenAndServe(cfg.MonitorAddr); err != nil {
				log.Error("monitor listen error", "monitor_addr", cfg.MonitorAddr, "error", err)
			}
		}()
	}
	// 逻辑代码部分
	GlobalDB, err = db.New(cfg.DB)
	if err != nil {
		log.Error("status db open error", "error", err)
		return
	}
	defer GlobalDB.Close()
	// 启动http server
	go func() {
		if err := http.ListenAndServe(cfg.BindAddr, nil); err != nil {
			log.Error("server start error", "error", err)
		}
	}()
	// 等待信号
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGUSR1)
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGUSR1:
			log.Info("melon got SIGUSR1 signal, change log")
			log.Init(cfg.LogFile, cfg.ErrFile, cfg.LogLevel)
			continue
		}
		monitor.Shutdown(context.Background())
		log.Info("melon got signal, exit", "sig", sig)
		break
	}
}
