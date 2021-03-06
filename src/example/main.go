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
	"nkwangwenfang.com/utils/pidfile"

	"example/router"
)

func main() {
	flag.Parse()
	// 解析配置文件
	cfg, err := loadConfig()
	if err != nil {
		return
	}
	// 设置日志
	log.Init(cfg.Log)
	// 创建pid文件
	if cfg.PidFile != "" {
		if err := pidfile.Save(cfg.PidFile); err != nil {
			log.Error("save pid file error", "pid_file", cfg.PidFile)
		}
		defer pidfile.Remove(cfg.PidFile)
	}
	// 启动monitor
	monitorServer := monitor.New(cfg.Monitor)
	defer monitorServer.Shutdown(context.Background())

	// 开始特定于http server的服务
	// 初始化路由设置
	router.Init()
	// 启动http server
	go func() {
		log.Debug("start http server", "addr", cfg.BindAddr)
		if err := http.ListenAndServe(cfg.BindAddr, nil); err != nil {
			log.Error("http server start error", "error", err)
		}
	}()

	// 等待信号
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGUSR1)
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGUSR1:
			log.Info("example got SIGUSR1 signal, change log")
			log.Init(cfg.Log)
			continue
		}
		log.Info("example got signal, exit", "sig", sig)
		break
	}
}
